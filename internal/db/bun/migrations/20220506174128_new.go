package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	l := logger.WithField("migration", "20220506174128")

	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			modelList := []interface{}{}
			for _, i := range modelList {
				l.Infof("creating table %T", i)
				if _, err := tx.NewCreateTable().Model(i).IfNotExists().Exec(ctx); err != nil {
					l.Errorf("can't create table %T: %s", i, err.Error())

					return err
				}
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
