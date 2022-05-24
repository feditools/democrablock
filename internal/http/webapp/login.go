package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/gorilla/sessions"
	nethttp "net/http"
)

// LoginGetHandler serves the login page.
func (m *Module) LoginGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	// init session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session) //nolint

	err := m.oauth.HandleLogin(w, r, us)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
	}
}
