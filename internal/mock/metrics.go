package mock

import (
	"github.com/feditools/democrablock/internal/metrics"
	"time"
)

// DBQuery is a new database query metric measurer.
type DBQuery struct{}

// Done is called when the db query is complete.
func (DBQuery) Done(_ bool) {}

// DBCacheQuery is a new database cache query metric measurer.
type DBCacheQuery struct{}

// Done is called when the db cache query is complete.
func (DBCacheQuery) Done(_, _ bool) {}

// MetricsCollector is a mock metrics collection.
type MetricsCollector struct{}

// Close does nothing.
func (MetricsCollector) Close() error {
	return nil
}

// DBQuery does nothing.
func (MetricsCollector) DBQuery(_ time.Duration, _ string, _ bool) {}

// NewDBQuery creates a new db query metrics collector.
func (MetricsCollector) NewDBQuery(_ string) metrics.DBQuery {
	return &DBQuery{}
}

// NewDBCacheQuery creates a new db cache query metrics collector.
func (MetricsCollector) NewDBCacheQuery(_ string) metrics.DBCacheQuery {
	return &DBCacheQuery{}
}

// HTTPRequestTiming does nothing.
func (MetricsCollector) HTTPRequestTiming(_ time.Duration, _ int, _, _ string) {}

// NewMetricsCollector creates a new mock metrics collector.
func NewMetricsCollector() (metrics.Collector, error) {
	return &MetricsCollector{}, nil
}
