package immudb

import (
	"context"
	immudb "github.com/codenotary/immudb/pkg/client"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/go-lib/metrics"
	"github.com/spf13/viper"
)

// New creates a new bun database immudb.
func New(ctx context.Context, m metrics.Collector) (db.DB, error) {
	l := logger.WithField("func", "New")

	opts := immudb.DefaultOptions().
		WithAddress(viper.GetString(config.Keys.DBAddress)).
		WithPort(viper.GetInt(config.Keys.DBPort))

	client := immudb.NewClient().WithOptions(opts)
	err := client.OpenSession(
		ctx,
		[]byte(viper.GetString(config.Keys.DBUser)),
		[]byte(viper.GetString(config.Keys.DBPassword)),
		viper.GetString(config.Keys.DBDatabase),
	)
	if err != nil {
		l.Errorf("opening session: %s", err.Error())

		return nil, err
	}

	return &Client{
		db:      client,
		metrics: m,
	}, nil
}

// Client is a DB interface compatible client for ImmuDB.
type Client struct {
	db      immudb.ImmuClient
	metrics metrics.Collector
}

func (c *Client) Close(ctx context.Context) db.Error {
	return c.db.CloseSession(ctx)
}

// ProcessError replaces any known values with our own db.Error types.
func (c *Client) ProcessError(err error) db.Error {
	switch {
	case err == nil:
		return nil
	default:
		return err
	}
}
