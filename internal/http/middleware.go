package http

import (
	nethttp "net/http"
	"time"
)

// MiddlewareMetrics sends http request metrics.
func (s *Server) MiddlewareMetrics(next nethttp.Handler) nethttp.Handler {
	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		start := time.Now()
		l := logger.WithField("func", "MiddlewareMetrics")

		wx := NewResponseWriter(w)

		// Do Request
		next.ServeHTTP(wx, r)

		ended := time.Since(start)
		l.Debugf("rendering %s took %d ms", r.URL.Path, ended.Milliseconds())
		go s.metrics.HTTPRequestTiming(ended, wx.Status(), r.Method, r.URL.Path)
	})
}
