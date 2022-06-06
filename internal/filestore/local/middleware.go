package local

import (
	"errors"
	nethttp "net/http"
	"strings"

	"github.com/feditools/democrablock/internal/kv"
	"github.com/feditools/democrablock/internal/path"
)

// middleware runs on every http request.
func (m *Module) middleware(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		l := logger.WithField("func", "middleware")

		// get token from query
		token, ok := r.URL.Query()[QueryToken]
		if !ok {
			returnErrorPage(w, nethttp.StatusUnauthorized)

			return
		}

		// check kv for token
		expectedObjectPath, err := m.kv.GetFileStorePresignedURL(r.Context(), token[0])
		if errors.Is(err, kv.ErrNil) {
			returnErrorPage(w, nethttp.StatusForbidden)

			return
		}
		if err != nil {
			l.Errorf("fs: %s", err.Error())
			returnErrorPage(w, nethttp.StatusInternalServerError)

			return
		}

		// check if token is for this path
		if strings.TrimPrefix(r.URL.Path, path.Filestore+"/") != expectedObjectPath {
			returnErrorPage(w, nethttp.StatusUnauthorized)

			return
		}

		// Do Request
		next.ServeHTTP(w, r)
	})
}
