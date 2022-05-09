package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/web"
	iofs "io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server
func (m *Module) Route(s *http.Server) error {
	staticFS, err := iofs.Sub(web.Files, DirStatic)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(staticFS))))

	//webapp := s.PathPrefix("/").Subrouter()
	return nil
}
