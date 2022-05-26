package webapp

import (
	nethttp "net/http"

	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/go-lib/language"
)

// HomeGetHandler serves the home page.
func (m *Module) HomeGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "HomeGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) //nolint

	// Init template variables
	tmplVars := &template.Home{}
	err := m.initTemplate(w, r, tmplVars)
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)

		return
	}
	err = m.initTemplatePublic(w, r, tmplVars)
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)

		return
	}

	tmplVars.PageTitle = localizer.TextDemocrablock().String()

	err = m.executeTemplate(w, template.HomeName, tmplVars)
	if err != nil {
		l.Errorf("could not render '%s' template: %s", template.HomeName, err.Error())
	}
}
