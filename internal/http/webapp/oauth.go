package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/login/pkg/oauth"
	"github.com/gorilla/sessions"
	nethttp "net/http"
)

// CallbackOauthGetHandler serves the home page.
func (m *Module) CallbackOauthGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "CallbackOauthGetHandler")

	// get localizer
	// localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) //nolint

	// get session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session) //nolint
	sessionID, ok := us.Values[SessionKeyID].(string)
	if !ok {
		l.Warn("missing session id")
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, "missing session id")

		return
	}

	token, err := m.oauth.HandleCallback(w, r, us, sessionID)
	if err != nil {
		if oerr, ok := err.(*oauth.Error); ok {
			m.returnErrorPage(w, r, oerr.Code, oerr.Message)

			return
		} else {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}
	}

	l.Debugf("login success: %+v", token)

	us.Values[SessionKeyOAuthToken] = token
	err = us.Save(r, w)
	if err != nil {
		l.Errorf("session oauth token: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	nethttp.Redirect(w, r, "/", nethttp.StatusFound)
}
