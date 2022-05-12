package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/go-lib/language"
	nethttp "net/http"
)

// MiddlewareRequireAdmin will redirect a user to login page if user not in context and will return unauthorized for
// a non admin user.
func (m *Module) MiddlewareRequireAdmin(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		account, shouldReturn := m.authRequireLoggedIn(w, r)
		if shouldReturn {
			return
		}

		if !account.Admin {
			localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) //nolint
			m.returnErrorPage(w, r, nethttp.StatusUnauthorized, localizer.TextUnauthorized().String())

			return
		}

		next.ServeHTTP(w, r)
	})
}
