package council

import (
	"context"
	"errors"

	"github.com/feditools/democrablock/cmd/democrablock/action"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/db/bun"
	"github.com/feditools/democrablock/internal/fedi"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/kv/redis"
	"github.com/feditools/democrablock/internal/models"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/go-lib"
	"github.com/feditools/go-lib/metrics/statsd"
	"github.com/spf13/viper"
)

// Init creates.
var Init action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Init")

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
		l.Info("closing metrics collector")
		err := metricsCollector.Close()
		if err != nil {
			l.Errorf("closing metrics: %s", err.Error())
		}
	}()

	// create database client
	l.Info("creating database client")
	dbClient, err := bun.New(ctx, metricsCollector)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	defer func() {
		l.Info("closing database client")
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

	// create fedi module
	fediMod, err := fedi.New(dbClient, httpClient, redisClient, tokz)
	if err != nil {
		l.Errorf("fedi: %s", err.Error())

		return err
	}

	// check if council exists
	councilCount, err := dbClient.CountFediAccountsWithCouncil(ctx)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	if councilCount > 0 {
		return errors.New("council already exists")
	}

	// check number of accounts, error if less than three
	newCouncil := viper.GetStringSlice(config.Keys.CouncilMembers)
	if len(newCouncil) < 3 {
		return errors.New("council must be at least three people")
	}

	// retrieve accounts if we don't have them
	members := make([]*models.FediAccount, len(newCouncil))
	for i, member := range newCouncil {
		username, domain, err := lib.SplitAccount(member)
		if err != nil {
			return err
		}

		// get instance
		instance, err := dbClient.ReadFediInstanceByDomain(ctx, domain)
		if err != nil {
			if !errors.Is(err, db.ErrNoEntries) {
				return err
			}

			newInstance, err := fediMod.NewFediInstanceFromDomain(ctx, domain)
			if err != nil {
				return err
			}

			err = dbClient.CreateFediInstance(ctx, newInstance)
			if err != nil {
				return err
			}

			instance = newInstance
		}

		// get account
		account, err := dbClient.ReadFediAccountByUsername(ctx, instance.ID, username)
		if err != nil {
			if !errors.Is(err, db.ErrNoEntries) {
				return err
			}

			newAccount, err := fediMod.NewFediAccountFromUsername(ctx, username, instance)
			if err != nil {
				return err
			}

			err = dbClient.CreateFediAccount(ctx, newAccount)
			if err != nil {
				return err
			}

			account = newAccount
		}

		members[i] = account
	}

	return nil
}
