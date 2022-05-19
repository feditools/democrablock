package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	nethttp "net/http"
)

// CallbackOauthGetHandler serves the home page.
func (m *Module) CallbackOauthGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "CallbackOauthGetHandler")

	// get localizer
	// localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) //nolint

	// get session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session) //nolint
	expectedState, ok := us.Values[SessionKeyOAuthState].(string)
	if !ok {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "missing state")

		return
	}
	expectedCode, ok := us.Values[SessionKeyOAuthCode].(string)
	if !ok {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "missing state")

		return
	}

	// delete so code and state can't be reused
	us.Values[SessionKeyOAuthState] = nil
	us.Values[SessionKeyOAuthCode] = nil
	err := us.Save(r, w)
	if err != nil {
		l.Errorf("session clearing oauth: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	if err := r.ParseForm(); err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}
	if state := r.Form.Get("state"); state != expectedState {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "State invalid")

		return
	}
	code := r.Form.Get("code")
	if code == "" {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "Code not found")

		return
	}

	l.Debugf("code: %s", code)

	token, err := m.oauth.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", expectedCode))
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	l.Debugf("login success: (%s), %+v", token.Type(), token)

	us.Values[SessionKeyOAuthToken] = token
	err = us.Save(r, w)
	if err != nil {
		l.Errorf("session oauth token: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	nethttp.Redirect(w, r, "/", nethttp.StatusFound)
}
