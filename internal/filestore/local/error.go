package local

import (
	"fmt"
	"net/http"

	"github.com/feditools/democrablock/internal/filestore"

	"github.com/tyrm/go-util/mimetype"
)

// ProcessError replaces any known values with our own db.Error types.
func (m *Module) ProcessError(err error) filestore.Error {
	switch {
	case err == nil:
		return nil
	default:
		return err
	}
}

func returnErrorPage(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", mimetype.TextPlain)
	_, err := w.Write([]byte(fmt.Sprintf("%d %s", code, http.StatusText(code))))
	if err != nil {
		logger.WithField("func", "returnErrorPage").Errorf("writing response: %s", err.Error())
	}
}

func (m *Module) methodNotAllowedHandler() http.Handler {
	// wrap in middleware since middleware isn't run on error pages
	return m.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		returnErrorPage(w, http.StatusMethodNotAllowed)
	}))
}

func (m *Module) notFoundHandler() http.Handler {
	// wrap in middleware since middleware isn't run on error pages
	return m.WrapInMiddlewares(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		returnErrorPage(w, http.StatusNotFound)
	}))
}

func (m *Module) WrapInMiddlewares(h http.Handler) http.Handler {
	return m.srv.WrapInMiddlewares(
		m.middleware(
			h,
		),
	)
}
