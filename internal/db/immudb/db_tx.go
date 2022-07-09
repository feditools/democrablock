package immudb

import (
	"context"

	"github.com/feditools/democrablock/internal/db"
)

func (*Client) TxNew(ctx context.Context) (db.TxID, db.Error) {
	// TODO implement me
	panic("implement me")
}

func (*Client) TxCommit(ctx context.Context, id db.TxID) db.Error {
	// TODO implement me
	panic("implement me")
}

func (*Client) TxRollback(ctx context.Context, id db.TxID) db.Error {
	// TODO implement me
	panic("implement me")
}
