package http

import (
	"context"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/metrics"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/tyrm/go-util/middleware"
	"net/http"
	"time"
)

const httpServerTimeout = 60 * time.Second

// Server is a http 2 web server.
type Server struct {
	metrics metrics.Collector
	router  *mux.Router
	srv     *http.Server
}

// NewServer creates a new http web server.
func NewServer(_ context.Context, m metrics.Collector) (*Server, error) {
	r := mux.NewRouter()

	s := &http.Server{
		Addr:         viper.GetString(config.Keys.ServerHTTPBind),
		Handler:      r,
		WriteTimeout: httpServerTimeout,
		ReadTimeout:  httpServerTimeout,
	}

	server := &Server{
		metrics: m,
		router:  r,
		srv:     s,
	}

	// add global middlewares
	r.Use(server.MiddlewareMetrics)
	r.Use(handlers.CompressHandler)
	r.Use(middleware.BlockFlocMux)

	return server, nil
}

// HandleFunc attaches a function to a path.
func (s *Server) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f)
}

// PathPrefix attaches a new route url path prefix.
func (s *Server) PathPrefix(path string) *mux.Route {
	return s.router.PathPrefix(path)
}

// Start starts the web server.
func (s *Server) Start() error {
	l := logger.WithField("func", "Start")
	l.Infof("listening on %s", s.srv.Addr)

	return s.srv.ListenAndServe()
}

// Stop shuts down the web server.
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
