package statsd

import (
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/democrablock/internal/metrics"
	"strconv"
	"time"
)

// HTTPRequestTiming send a metrics relating to a http request.
func (m *Module) HTTPRequestTiming(t time.Duration, status int, method, path string) {
	err := m.s.TimingDuration(
		metrics.StatHTTPRequest,
		t,
		m.rate,
		statsd.Tag{metrics.TagHTTPStatus, strconv.Itoa(status)},
		statsd.Tag{metrics.TagHTTPMethod, method},
		statsd.Tag{metrics.TagHTTPPath, path},
	)
	if err != nil {
		logger.WithField("func", "HTTPRequestTiming").Warn(err.Error())
	}
}
