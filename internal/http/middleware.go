package http

import (
	"net/http"

	"github.com/go-http-utils/etag"
	"github.com/gorilla/handlers"
	"github.com/tyrm/go-util/middleware"
)

// MiddlewareMetrics sends http request metrics.
func (s *Server) MiddlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric := s.metrics.NewHTTPRequest(r.Method, r.URL.Path)
		l := logger.WithField("func", "middlewareMetrics")

		wx := NewResponseWriter(w)

		// Do Request
		next.ServeHTTP(wx, r)

		go func() {
			ended := metric.Done(wx.Status())
			l.Debugf("rendering %s took %d ms", r.URL.Path, ended.Milliseconds())
		}()
	})
}

// WrapInMiddlewares wraps an http.Handler in the server's middleware.
func (s *Server) WrapInMiddlewares(h http.Handler) http.Handler {
	return s.MiddlewareMetrics(
		middleware.BlockMissingUserAgentMux(
			etag.Handler(
				handlers.CompressHandler(
					middleware.BlockFlocMux(
						h,
					),
				), false,
			),
		),
	)
}
