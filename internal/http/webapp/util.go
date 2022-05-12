package webapp

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/models"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/web"
	"github.com/gorilla/sessions"
	"io/ioutil"
	nethttp "net/http"
)

// auth helpers

func (m *Module) authRequireLoggedIn(w nethttp.ResponseWriter, r *nethttp.Request) (*models.FediAccount, bool) {
	us := r.Context().Value(http.ContextKeySession).(*sessions.Session)

	if r.Context().Value(http.ContextKeyAccount) == nil {
		// Save current page
		if r.URL.Query().Encode() == "" {
			us.Values[SessionKeyLoginRedirect] = r.URL.Path
		} else {
			us.Values[SessionKeyLoginRedirect] = r.URL.Path + "?" + r.URL.Query().Encode()
		}
		err := us.Save(r, w)
		if err != nil {
			m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())
			return nil, true
		}

		// redirect to login
		nethttp.Redirect(w, r, path.Login, nethttp.StatusFound)
		return nil, true
	}

	account := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount)
	return account, false
}

// signature caching

func getSignature(path string) (string, error) {
	l := logger.WithField("func", "getSignature")

	file, err := web.Files.Open(path)
	if err != nil {
		l.Errorf("opening file: %s", err.Error())

		return "", err
	}

	// read it
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// hash it
	h := sha512.New384()
	_, err = h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("sha384-%s", signature), nil
}

func (m *Module) getSignatureCached(path string) (string, error) {
	if sig, ok := m.readCachedSignature(path); ok {
		return sig, nil
	}
	sig, err := getSignature(path)
	if err != nil {
		return "", err
	}
	m.writeCachedSignature(path, sig)

	return sig, nil
}

func (m *Module) readCachedSignature(path string) (string, bool) {
	m.sigCacheLock.RLock()
	val, ok := m.sigCache[path]
	m.sigCacheLock.RUnlock()

	return val, ok
}

func (m *Module) writeCachedSignature(path string, sig string) {
	m.sigCacheLock.Lock()
	m.sigCache[path] = sig
	m.sigCacheLock.Unlock()
}
