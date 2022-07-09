package immudb

import (
	"context"
	"github.com/feditools/democrablock/internal/db"
)

func (c *Client) TxNew(ctx context.Context) (db.TxID, db.Error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TxCommit(ctx context.Context, id db.TxID) db.Error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) TxRollback(ctx context.Context, id db.TxID) db.Error {
	//TODO implement me
	panic("implement me")
}
