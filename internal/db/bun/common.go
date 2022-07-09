package bun

import (
	"context"
	"errors"
	"fmt"
	"github.com/uptrace/bun"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/db/bun/migrations"
	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/migrate"
)

// Close closes the bun db connection.
func (c *Client) Close(_ context.Context) db.Error {
	l := logger.WithField("func", "Close")
	l.Info("closing db connection")

	return c.bun.Close()
}

// DoMigration runs schema migrations on the database.
func (c *Client) DoMigration(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	migrator := migrate.NewMigrator(c.bun.DB, migrations.Migrations)

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

// LoadTestData adds test data to the database.
func (c *Client) LoadTestData(ctx context.Context) db.Error {
	l := logger.WithField("func", "LoadTestData")
	l.Debugf("adding test data")

	// Truncate
	modelList := []interface{}{}

	for _, m := range modelList {
		l.Debugf("truncating %T", m)
		_, err := c.bun.NewTruncateTable().Model(m).Exec(ctx)
		if err != nil {
			l.Errorf("truncating %T: %s", m, err.Error())

			return err
		}
	}

	// fix sequences
	sequences := []struct {
		table        string
		currentValue int
	}{}

	switch c.bun.Dialect().Name() {
	case dialect.Invalid:
		return fmt.Errorf("invalid dialect %s", c.bun.Dialect().Name())
	case dialect.MSSQL:
		return errors.New("dialect MSSQL unsupported")
	case dialect.MySQL:
		return errors.New("dialect MySQL unsupported")
	case dialect.SQLite:
		// nothing to do
	case dialect.PG:
		for _, s := range sequences {
			_, err := c.bun.Exec("SELECT setval(?, ?, true);", fmt.Sprintf("%s_id_seq", s.table), s.currentValue)
			if err != nil {
				l.Errorf("can't update sequence for %s: %s", s.table, err.Error())

				return err
			}
		}
	default:
		return fmt.Errorf("unknown dialect %s", c.bun.Dialect().Name())
	}

	return nil
}

// ResetCache does nothing. This module contains no cache.
func (*Client) ResetCache(_ context.Context) db.Error {
	return nil
}

func create(ctx context.Context, c bun.IDB, i interface{}) error {
	_, err := c.NewInsert().Model(i).Exec(ctx)
	if err != nil {
		logger.WithField("func", "create").Errorf("db: %s", err.Error())
	}

	return err
}

func update(ctx context.Context, c bun.IDB, i interface{}) error {
	q := c.NewUpdate().Model(i).WherePK()

	_, err := q.Exec(ctx)

	return err
}
