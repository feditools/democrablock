package cachemem

import (
	"context"

	"github.com/feditools/democrablock/internal/db"
)

func (c *CacheMem) TxNew(ctx context.Context) (db.TxID, db.Error) {
	return c.db.TxNew(ctx)
}

func (c *CacheMem) TxCommit(ctx context.Context, id db.TxID) db.Error {
	return c.db.TxCommit(ctx, id)
}

func (c *CacheMem) TxRollback(ctx context.Context, id db.TxID) db.Error {
	return c.db.TxRollback(ctx, id)
}
