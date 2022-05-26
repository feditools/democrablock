package webapp

import (
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/web"
	"io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server.
func (m *Module) Route(s *http.Server) error {
	staticFS, err := fs.Sub(web.Files, DirStatic)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(staticFS))))

	webapp := s.PathPrefix("/").Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(path.CallbackOauth, m.CallbackOauthGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.Home, m.HomeGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.Login, m.LoginGetHandler).Methods(nethttp.MethodGet)
	webapp.HandleFunc(path.Login, m.LoginPostHandler).Methods(nethttp.MethodPost)
	webapp.HandleFunc(path.Logout, m.LogoutGetHandler).Methods(nethttp.MethodGet)

	return nil
}
