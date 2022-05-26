package webapp

import (
	nethttp "net/http"

	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/go-lib/language"
)

func makePublicNavbar(r *nethttp.Request) template.Navbar {
	// get localizer
	localizer := r.Context().Value(http.ContextKeyLocalizer).(*language.Localizer) //nolint

	// create navbar
	newNavbar := template.Navbar{
		{
			Text:     localizer.TextHomeWeb().String(),
			MatchStr: path.ReHome,
			FAIcon:   "home",
			URL:      path.Home,
		},
		{
			Text:     localizer.TextList(1).String(),
			MatchStr: path.ReList,
			FAIcon:   "bars-staggered",
			URL:      path.List,
		},
	}

	newNavbar.ActivateFromPath(r.URL.Path)

	return newNavbar
}
