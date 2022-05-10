package statsd

import (
	"sync"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/metrics"
	"github.com/spf13/viper"
)

const DefaultRate = 1.0

// Module represents a statsd metrics collector.
type Module struct {
	s statsd.Statter

	rate                 float32
	systemCollectionOnce sync.Once
	systemCollectionRate time.Duration

	done chan bool
}

// New creates a new Statsd metrics module.
func New() (metrics.Collector, error) {
	statsConfig := &statsd.ClientConfig{
		Address: viper.GetString(config.Keys.MetricsStatsDAddress),
		Prefix:  viper.GetString(config.Keys.MetricsStatsDPrefix),
	}
	client, err := statsd.NewClientWithConfig(statsConfig)
	if err != nil {
		return nil, err
	}

	m := &Module{
		s: client,

		rate:                 DefaultRate,
		systemCollectionRate: 10 * time.Second,

		done: make(chan bool),
	}

	m.systemCollectionOnce.Do(m.systemCollector)

	return m, nil
}

// Close closes the statsd metrics collector.
func (m *Module) Close() error {
	return m.s.Close()
}
