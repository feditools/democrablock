package immudb

import (
	"context"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/db/immudb/migrations"
	migrate "github.com/tyrm/immudb-migrate"
)

func (c *Client) DoMigration(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	migrator := migrate.NewMigrator(c.db, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		if err.Error() == "migrate: there are no any migrations" {
			return nil
		}

		return err
	}

	if group.ID == 0 {
		l.Info("there are no new migrations to run")

		return nil
	}

	l.Infof("migrated database to %s", group)

	return nil
}

func (c *Client) LoadTestData(ctx context.Context) db.Error {
	// TODO implement me
	panic("implement me")
}

func (c *Client) ResetCache(_ context.Context) db.Error {
	return nil
}
