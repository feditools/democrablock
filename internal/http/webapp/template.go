package webapp

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/democrablock/internal/models"
	"github.com/feditools/go-lib/language"
	nethttp "net/http"
)

func (m *Module) initTemplate(_ nethttp.ResponseWriter, r *nethttp.Request, tmpl template.InitTemplate) error {
	// set text handler
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) //nolint
	tmpl.SetLocalizer(localizer)

	// set language
	if lang, ok := r.Context().Value(http.ContextKeyLanguage).(string); ok {
		tmpl.SetLanguage(lang)
	}

	// set logo image src
	tmpl.SetLogoSrc(m.logoSrcDark, m.logoSrcLight)

	// add css
	for _, link := range m.headLinks {
		tmpl.AddHeadLink(link)
	}

	// add scripts
	for _, script := range m.footerScripts {
		tmpl.AddFooterScript(script)
	}

	if account, ok := r.Context().Value(http.ContextKeyAccount).(*models.FediAccount); ok {
		tmpl.SetAccount(account)
	}

	return nil
}

func (m *Module) executeTemplate(w nethttp.ResponseWriter, name string, tmplVars interface{}) error {
	b := new(bytes.Buffer)
	err := m.templates.ExecuteTemplate(b, name, tmplVars)
	if err != nil {
		return err
	}

	h := sha256.New()

	_, err = h.Write(b.Bytes())
	if err != nil {
		return err
	}
	w.Header().Set("Digest", fmt.Sprintf("sha-256=%s", base64.StdEncoding.EncodeToString(h.Sum(nil))))

	if m.minify == nil {
		_, err := w.Write(b.Bytes())

		return err
	}

	return m.minify.Minify("text/html", w, b)
}
