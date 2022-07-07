package local

import (
	"encoding/hex"
	nethttp "net/http"
	"strings"

	"github.com/feditools/democrablock/internal/path"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/gorilla/mux"
)

func (m *Module) handleGet(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "handleGet")
	l.Debugf("boop")

	// get vars
	vars := mux.Vars(r)

	group, gok := vars[path.VarGroupID]
	hash1, h1ok := vars[path.VarHash1ID]
	hash2, h2ok := vars[path.VarHash2ID]
	hash3, h3ok := vars[path.VarHash3ID]
	fshash, fshok := vars[path.VarFileStoreHashID]
	fssuffix, fssok := vars[path.VarFileStoreSuffixID]

	if !gok || !h1ok || !h2ok || !h3ok || !fshok || !fssok {
		l.Debugf(
			"missing var: group(%t), hash1(%t), hash2(%t), hash3(%t), fshash(%t), fssuffix(%t)",
			gok,
			h1ok,
			h2ok,
			h3ok,
			fshok,
			fssok)
		returnErrorPage(w, nethttp.StatusBadRequest)

		return
	}

	if !checkPathSanity(fshash, hash1, hash2, hash3) {
		l.Debugf(
			"insane path: %s",
			strings.TrimPrefix(r.URL.Path, path.Filestore+"/"),
		)
		returnErrorPage(w, nethttp.StatusBadRequest)

		return
	}

	// get file
	hash, err := hex.DecodeString(fshash)
	if err != nil {
		l.Errorf(
			"decoding hash '%s': %s",
			fshash,
			err.Error(),
		)
		returnErrorPage(w, nethttp.StatusInternalServerError)

		return
	}
	data, err := m.GetFile(r.Context(), group, hash, fssuffix)
	if err != nil {
		l.Errorf(
			"reading file '%s': %s",
			fshash,
			err.Error(),
		)
		returnErrorPage(w, nethttp.StatusInternalServerError)

		return
	}

	_, err = w.Write(data)
	if err != nil {
		l.Errorf(
			"writing response '%s': %s",
			fshash,
			err.Error(),
		)
		returnErrorPage(w, nethttp.StatusInternalServerError)

		return
	}

	w.WriteHeader(nethttp.StatusOK)
	w.Header().Set("Content-Type", string(libhttp.ToMime(libhttp.Suffix(fssuffix))))
}

func checkPathSanity(hash string, bits ...string) bool {
	// check bit count
	if len(bits) != 3 {
		return false
	}

	// generate prefix
	var prefix string
	for _, b := range bits {
		prefix += b
	}

	// check for prefix
	return strings.HasPrefix(hash, prefix)
}
