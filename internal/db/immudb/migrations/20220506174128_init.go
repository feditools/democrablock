package migrations

import (
	"context"

	"github.com/codenotary/immudb/pkg/client"
	statements "github.com/feditools/democrablock/internal/db/immudb/migrations/20220506174128_init"
)

func init() {
	l := logger.WithField("migration", "20220506174128")

	up := func(ctx context.Context, tx client.Tx) error {
		l.Debugf("running up migration")

		// run statements
		creates := []string{
			statements.CreateTableFediInstances,
			statements.CreateIndexFediInstancesUnique,
			statements.CreateTableFediAccounts,
			statements.CreateIndexFediAccountsUnique,
		}
		for _, c := range creates {
			if err := tx.SQLExec(ctx, c, nil); err != nil {
				l.Errorf("SQLExec: %s", err.Error())

				return err
			}
		}

		return nil
	}

	if err := Migrations.Register(up); err != nil {
		panic(err)
	}
}
