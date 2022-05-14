package webapp

import (
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/web"
	iofs "io/fs"
	nethttp "net/http"
)

// Route attaches routes to the web server.
func (m *Module) Route() error {
	staticFS, err := iofs.Sub(web.Files, DirStatic)
	if err != nil {
		return err
	}

	// Static Files
	m.srv.PathPrefix(path.Static).Handler(nethttp.StripPrefix(path.Static, nethttp.FileServer(nethttp.FS(staticFS))))

	webapp := m.srv.PathPrefix("/").Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(path.Home, m.HomeGetHandler).Methods("GET")
	webapp.HandleFunc(path.Login, m.LoginGetHandler).Methods("GET")
	webapp.HandleFunc(path.CallbackOauth, m.CallbackOauthGetHandler).Methods("GET")

	return nil
}
