package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/gorilla/sessions"
	nethttp "net/http"
)

// LoginGetHandler serves the login page.
func (m *Module) LoginGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "LoginGetHandler")

	// init session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session) //nolint
	sessionID, ok := us.Values[SessionKeyID].(string)
	if !ok {
		l.Warn("missing session id")
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, "missing session id")

		return
	}

	err := m.oauth.HandleLogin(w, r, us, sessionID)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
	}
}
