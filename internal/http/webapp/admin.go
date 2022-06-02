package webapp

import (
	nethttp "net/http"

	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/go-lib/language"
)

func makeAdminNavbar(r *nethttp.Request) template.Navbar {
	// get localizer
	l := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) // nolint

	// create navbar
	newNavbar := template.Navbar{
		{
			Text:     l.TextHomeWeb().String(),
			MatchStr: path.ReAdmin,
			FAIcon:   "home",
			URL:      path.Admin,
		},
		{
			Text:     l.TextFediverse().String(),
			MatchStr: path.ReAdminFediversePre,
			FAIcon:   "home",
			URL:      path.AdminFediverse,
		},
		{
			Text:     l.TextSystem(1).String(),
			MatchStr: path.ReAdminSystemPre,
			FAIcon:   "desktop",
			URL:      path.AdminSystem,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
