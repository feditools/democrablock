package server

import (
	"context"
	"fmt"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/grpc"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/webapp"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/go-lib/language"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"

	"github.com/feditools/democrablock/cmd/democrablock/action"
	"github.com/feditools/democrablock/internal/db/bun"
	cachemem "github.com/feditools/democrablock/internal/db/cache_mem"
	"github.com/feditools/democrablock/internal/kv/redis"
	"github.com/feditools/democrablock/internal/metrics/statsd"
	"github.com/tyrm/go-util"
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

	grcpClient, err := grpc.New(ctx)
	if err != nil {
		l.Errorf("grpc: %s", err.Error())

		return err
	}

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

	tokz, err := token.New()
	if err != nil {
		l.Errorf("create tokenizer: %s", err.Error())

		return err
	}

	languageMod, err := language.New()
	if err != nil {
		l.Errorf("language: %s", err.Error())

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
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleWebapp) {
		l.Infof("adding webapp module")
		webMod, err := webapp.New(ctx, cachedDBClient, grcpClient, redisClient, languageMod, tokz, metricsCollector)
		if err != nil {
			logrus.Errorf("webapp module: %s", err.Error())

			return err
		}
		webModules = append(webModules, webMod)
	}

	// add modules to server
	for _, mod := range webModules {
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
