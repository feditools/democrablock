package local

import (
	nethttp "net/http"
)

// Middleware runs on every http request.
func (m *Module) Middleware(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		l := logger.WithField("func", "Middleware")

		token, ok := r.URL.Query()[QueryToken]
		if !ok {
			returnErrorPage(w, nethttp.StatusUnauthorized)

			return
		}

		expectedObjectPath, err := m.kv.GetFileStorePresignedURL(r.Context(), token[0])
		if err != nil {
			l.Errorf("fs: %s", err.Error())
			returnErrorPage(w, nethttp.StatusInternalServerError)

			return
		}

		// get vars
		/*vars := mux.Vars(r)

		group, gok := vars[path.VarGroupID]
		hash1, h1ok := vars[path.VarHash1ID]
		hash2, h2ok := vars[path.VarHash2ID]
		hash3, h3ok := vars[path.VarHash2ID]
		fshash, fshok := vars[path.VarFileStoreHashID]
		fssuffix, fssok := vars[path.VarFileStoreSuffixID]

		if !gok || !h1ok || !h2ok || !h3ok || !fshok || !fssok {
			returnErrorPage(w, nethttp.StatusBadRequest)

			return
		}*/

		// reconstruct url
		l.Debugf("expected asdf: %s", expectedObjectPath)
		l.Debugf("asdf: %s", r.URL.Path)

		// Do Request
		next.ServeHTTP(w, r)
	})
}
