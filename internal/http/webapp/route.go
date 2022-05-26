package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/web"
	iofs "io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *http.Server) error {
	staticFS, err := iofs.Sub(web.Files, DirStatic)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(staticFS))))

	s.HandleFunc(path.CallbackOauth, m.CallbackOauthGetHandler).Methods(nethttp.MethodGet)
	s.HandleFunc(path.Login, m.LoginGetHandler).Methods(nethttp.MethodGet)
	s.HandleFunc(path.Login, m.LoginPostHandler).Methods(nethttp.MethodPost)
	s.HandleFunc(path.Logout, m.LogoutGetHandler).Methods(nethttp.MethodGet)

	// webapp := s.PathPrefix("/").Subrouter()
	return nil
}
