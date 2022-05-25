package webapp

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/feditools/democrablock/internal/http"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/go-http-utils/etag"
	"golang.org/x/oauth2"
	nethttp "net/http"
	"strconv"
)

// Middleware runs on every http request.
func (m *Module) Middleware(next nethttp.Handler) nethttp.Handler {
	return etag.Handler(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		l := logger.WithField("func", "Middleware")

		// create localizer
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer, err := m.language.NewLocalizer(lang, accept)
		if err != nil {
			l.Errorf("could get localizer: %s", err.Error())
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

			return
		}
		ctx := context.WithValue(r.Context(), http.ContextKeyLocalizer, localizer)

		// Init Session
		us, err := m.store.Get(r, "democrablock")
		if err != nil {
			l.Errorf("get session: %s", err.Error())
			m.returnErrorPage(w, r.WithContext(ctx), nethttp.StatusInternalServerError, err.Error())

			return
		}

		ctx = context.WithValue(ctx, http.ContextKeySession, us)

		// set request language
		ctx = context.WithValue(ctx, http.ContextKeyLanguage, libhttp.GetPageLang(lang, accept, m.language.Language().String()))

		// retrieve our token
		if token, ok := us.Values[SessionKeyOAuthToken].(oauth2.Token); ok {
			l.Debugf("token: %+v", token)
			newToken, newIDToken, refreshed, err := m.oauth.TokenSource(r.Context(), us, &token)
			if err != nil {
				l.Errorf("token source: %s", err.Error())
				m.returnErrorPage(w, r.WithContext(ctx), nethttp.StatusInternalServerError, err.Error())

				return
			}
			if refreshed {
				us.Values[SessionKeyOAuthToken] = newToken
				us.Values[SessionKeyOAuthJWT] = newIDToken

				if err := us.Save(r, w); err != nil {
					l.Errorf("save session: %s", err.Error())
					m.returnErrorPage(w, r.WithContext(ctx), nethttp.StatusInternalServerError, fmt.Sprintf("save session: %s", err.Error()))

					return
				}
			}
		}

		// Retrieve our account and type-assert it
		if accountID, ok := us.Values[SessionKeyOAuthJWT].(oidc.IDToken); ok {
			l.Debugf("accountID: %+v", accountID)

			// parse subject integer
			accountIDInt, err := strconv.ParseInt(accountID.Subject, 10, 64)
			if err != nil {
				l.Errorf("parsing int: %s", err.Error())
				m.returnErrorPage(w, r.WithContext(ctx), nethttp.StatusInternalServerError, err.Error())

				return
			}

			// read federated accounts
			account, err := m.grpc.GetFediAccount(ctx, accountIDInt)
			if err != nil {
				l.Errorf("db read: %s", err.Error())
				m.returnErrorPage(w, r.WithContext(ctx), nethttp.StatusInternalServerError, err.Error())

				return
			}

			if account != nil {
				// read federated instance
				instance, err := m.grpc.GetFediInstance(ctx, account.InstanceID)
				if err != nil {
					l.Errorf("db read: %s", err.Error())
					m.returnErrorPage(w, r.WithContext(ctx), nethttp.StatusInternalServerError, err.Error())

					return
				}
				account.Instance = instance

				ctx = context.WithValue(ctx, http.ContextKeyAccount, account)
			}
		}

		// Do Request
		next.ServeHTTP(w, r.WithContext(ctx))
	}), false)
}
