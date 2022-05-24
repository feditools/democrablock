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

	token, idToken, err := m.oauth.HandleCallback(w, r, us)
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
	us.Values[SessionKeyOAuthJWT] = idToken
	err = us.Save(r, w)
	if err != nil {
		l.Errorf("session oauth token: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	nethttp.Redirect(w, r, "/", nethttp.StatusFound)
}
