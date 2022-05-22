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
	sessionID, ok := us.Values[SessionKeyID].(string)
	if !ok {
		l.Warn("missing session id")
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, "missing session id")

		return
	}
	expectedCode, ok := us.Values[SessionKeyOAuthCode].(string)
	if !ok {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "missing state")

		return
	}
	expectedNonce, ok := us.Values[SessionKeyOAuthNonce].(string)
	if !ok {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "missing state")

		return
	}
	expectedState, ok := us.Values[SessionKeyOAuthState].(string)
	if !ok {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "missing state")

		return
	}

	// delete so code and state can't be reused
	us.Values[SessionKeyOAuthCode] = nil
	us.Values[SessionKeyOAuthState] = nil
	err := us.Save(r, w)
	if err != nil {
		l.Errorf("session clearing oauth: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	// parse form
	if err := r.ParseForm(); err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	// compare state
	if state := r.Form.Get("state"); state != expectedState {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "State invalid")

		return
	}

	// get code
	code := r.Form.Get("code")
	if code == "" {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "Code not found")

		return
	}

	// request token
	token, err := m.oauth.Exchange(
		r.Context(),
		code,
		oauth2.SetAuthURLParam("session_id", sessionID),
		oauth2.SetAuthURLParam("code_verifier", expectedCode),
	)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}
	l.Debugf("exchange: (%s), %+v", token.Type(), token)

	// validate token
	rawIDToken := token.AccessToken
	/*rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, "id_token field missing")

		return
	}*/
	l.Debugf("raw id token: %s", rawIDToken)
	idToken, err := m.oauthVerifier.Verify(r.Context(), rawIDToken)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}
	if idToken.Nonce != expectedNonce {
		m.returnErrorPage(w, r, nethttp.StatusBadRequest, "nonce did not match")

		return
	}

	l.Debugf("login success: %+v", idToken)

	us.Values[SessionKeyOAuthToken] = token
	err = us.Save(r, w)
	if err != nil {
		l.Errorf("session oauth token: %s", err.Error())
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	nethttp.Redirect(w, r, "/", nethttp.StatusFound)
}
