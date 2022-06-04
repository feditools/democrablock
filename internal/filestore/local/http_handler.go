package local

import nethttp "net/http"

func (m *Module) handleGet(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "handleGet")
	l.Debugf("boop")
}
