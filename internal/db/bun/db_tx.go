package bun

import (
	"context"
	"database/sql"

	"github.com/feditools/democrablock/internal/db"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (c *Client) TxNew(ctx context.Context) (db.TxID, db.Error) {
	id := db.TxID(uuid.New().String())

	tx, err := c.bun.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return "", err
	}
	c.setTx(id, &tx)

	return id, nil
}

func (c *Client) TxCommit(_ context.Context, id db.TxID) db.Error {
	tx, err := c.getTx(id)
	if err != nil {
		return err
	}
	c.deleteTx(id)

	return tx.Commit()
}

func (c *Client) TxRollback(_ context.Context, id db.TxID) db.Error {
	tx, err := c.getTx(id)
	if err != nil {
		return err
	}
	c.deleteTx(id)

	return tx.Rollback()
}

func (c *Client) deleteTx(id db.TxID) {
	c.txLock.Lock()
	defer c.txLock.Unlock()

	delete(c.tx, id)
}

func (c *Client) getTx(id db.TxID) (*bun.Tx, db.Error) {
	c.txLock.RLock()
	defer c.txLock.RUnlock()

	tx, ok := c.tx[id]
	if !ok {
		return nil, db.ErrUnknown
	}

	return tx, nil
}

func (c *Client) setTx(id db.TxID, tx *bun.Tx) {
	c.txLock.Lock()
	defer c.txLock.Unlock()

	c.tx[id] = tx
}
