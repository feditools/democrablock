package webapp

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/feditools/democrablock/internal/http"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	nethttp "net/http"
)

// LoginGetHandler serves the login page.
func (m *Module) LoginGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	// init session
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session) //nolint

	newState := uuid.New().String()
	newCode := uuid.New().String()
	us.Values[SessionKeyOAuthState] = newState
	us.Values[SessionKeyOAuthCode] = newCode
	if err := us.Save(r, w); err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	u := m.oauth.AuthCodeURL(
		newState,
		oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256(newCode)),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)

	nethttp.Redirect(w, r, u, nethttp.StatusFound)
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))

	return base64.URLEncoding.EncodeToString(s256[:])
}
