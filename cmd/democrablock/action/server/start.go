package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/feditools/democrablock/internal/db/immudb"

	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/fedi"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/webapp"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/go-lib/language"
	"github.com/spf13/viper"

	"github.com/feditools/democrablock/cmd/democrablock/action"
	"github.com/feditools/democrablock/internal/kv/redis"
	"github.com/feditools/go-lib"
	"github.com/feditools/go-lib/metrics/statsd"
)

// Start starts the server.
var Start action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Start")
	l.Infof("starting")

	// create metrics collector
	metricsCollector, err := statsd.New(
		viper.GetString(config.Keys.MetricsStatsDAddress),
		viper.GetString(config.Keys.MetricsStatsDPrefix),
	)
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

	// create db client
	dbClient, err := immudb.New(ctx, metricsCollector)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	// create kv client
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

	// create http client
	httpClient, err := http.NewClient(ctx)
	if err != nil {
		l.Errorf("http client: %s", err.Error())

		return err
	}

	// create tokenizer
	tokz, err := token.New()
	if err != nil {
		l.Errorf("create tokenizer: %s", err.Error())

		return err
	}

	// create language module
	languageMod, err := language.New()
	if err != nil {
		l.Errorf("language: %s", err.Error())

		return err
	}

	// create fedi module
	fediMod, err := fedi.New(dbClient, httpClient, redisClient, tokz)
	if err != nil {
		l.Errorf("fedi: %s", err.Error())

		return err
	}

	// create http server
	l.Debug("creating http server")
	httpServer, err := http.NewServer(ctx, metricsCollector)
	if err != nil {
		l.Errorf("http httpServer: %s", err.Error())

		return err
	}

	// create web modules
	var webModules []http.Module
	if lib.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleWebapp) {
		l.Infof("adding webapp module")
		webMod, err := webapp.New(ctx, dbClient, redisClient, fediMod, languageMod, tokz, metricsCollector)
		if err != nil {
			l.Errorf("webapp module: %s", err.Error())

			return err
		}
		webModules = append(webModules, webMod)
	}

	// add modules to server
	for _, mod := range webModules {
		mod.SetServer(httpServer)
		err := mod.Route(httpServer)
		if err != nil {
			l.Errorf("loading %s module: %s", mod.Name(), err.Error())

			return err
		}
	}

	// ** start application **
	errChan := make(chan error)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	stopSigChan := make(chan os.Signal, 1)
	signal.Notify(stopSigChan, syscall.SIGINT, syscall.SIGTERM)

	// start webserver
	go func(s *http.Server, errChan chan error) {
		l.Debug("starting http server")
		err := s.Start()
		if err != nil {
			errChan <- fmt.Errorf("http server: %s", err.Error())
		}
	}(httpServer, errChan)

	// wait for event
	select {
	case sig := <-stopSigChan:
		l.Infof("got sig: %s", sig)
	case err := <-errChan:
		l.Fatal(err.Error())
	}

	l.Infof("done")

	return nil
}
