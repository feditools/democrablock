package server

import (
	"context"

	"github.com/feditools/democrablock/cmd/login/action"
	"github.com/feditools/democrablock/internal/db/bun"
	cachemem "github.com/feditools/democrablock/internal/db/cache_mem"
	"github.com/feditools/democrablock/internal/kv/redis"
	"github.com/feditools/democrablock/internal/metrics/statsd"
)

// Start starts the server.
var Start action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Start")

	l.Infof("starting")
	metricsCollector, err := statsd.New()
	if err != nil {
		l.Errorf("metrics: %s", err.Error())

		return err
	}
	defer func() {
		err := metricsCollector.Close()
		if err != nil {
			l.Errorf("closing metrics: %s", err.Error())
		}
	}()

	dbClient, err := bun.New(ctx, metricsCollector)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	cachedDBClient, err := cachemem.New(ctx, dbClient, metricsCollector)
	if err != nil {
		l.Errorf("db-cachemem: %s", err.Error())

		return err
	}
	defer func() {
		err := cachedDBClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	redisClient, err := redis.New(ctx)
	if err != nil {
		l.Errorf("redis: %s", err.Error())

		return err
	}
	defer func() {
		err := redisClient.Close(ctx)
		if err != nil {
			l.Errorf("closing redis: %s", err.Error())
		}
	}()

	return nil
}
