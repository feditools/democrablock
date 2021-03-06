package webapp

import (
	nethttp "net/http"

	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
)

// AdminFediGetHandler serves the admin fediverse page.
func (m *Module) AdminFediGetHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "AdminFediverseGetHandler")

	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) // nolint

	// Init template variables
	tmplVars := &template.AdminFedi{
		Common: template.Common{
			PageTitle: localizer.TextFediverse().String(),
		},
		Admin: template.Admin{
			Sidebar: makeAdminFediverseSidebar(r),
		},
	}

	err := m.initTemplateAdmin(w, r, tmplVars)
	if err != nil {
		m.returnErrorPage(w, r, nethttp.StatusInternalServerError, err.Error())

		return
	}

	err = m.executeTemplate(w, template.AdminFediName, tmplVars)
	if err != nil {
		l.Errorf("could not render %s template: %s", template.AdminFediName, err.Error())
	}
}

func makeAdminFediverseSidebar(r *nethttp.Request) libtemplate.Sidebar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) // nolint

	// create sidebar
	newSidebar := libtemplate.Sidebar{
		{
			Text: localizer.TextOauth20Settings().String(),
			Children: []libtemplate.SidebarNode{
				{
					Text:    localizer.TextInstance(2).String(),
					Matcher: path.ReAdminFediverseInstancesPre,
					Icon:    "desktop",
					URI:     path.AdminFediverseInstances,
				},
				{
					Text:    localizer.TextAccount(2).String(),
					Matcher: path.ReAdminFediverseAccountsPre,
					Icon:    "user",
					URI:     path.AdminFediverseAccounts,
				},
			},
		},
	}

	newSidebar.ActivateFromPath(r.URL.Path)

	return newSidebar
}
