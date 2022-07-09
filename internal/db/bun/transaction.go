package bun

import (
	"context"
	"errors"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/models"
	libdatabase "github.com/feditools/go-lib/database"
	"github.com/uptrace/bun"
)

func (c *Client) CountTransactions(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountTransactions")

	count, err := newTransactionQ(c.bun, (*models.Transaction)(nil)).Count(ctx)
	if err != nil {
		go metric.Done(true)

		return 0, c.bun.errProc(err)
	}
	go metric.Done(false)

	return int64(count), nil
}

func (c *Client) CreateTransaction(ctx context.Context, transaction *models.Transaction) db.Error {
	metric := c.metrics.NewDBQuery("CreateTransaction")

	if err := create(ctx, c.bun, transaction); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}
	go metric.Done(false)

	return nil
}

func (c *Client) CreateTransactionTX(ctx context.Context, txID db.TxID, transaction *models.Transaction) db.Error {
	metric := c.metrics.NewDBQuery("CreateTransactionTX")

	tx, err := c.getTx(txID)
	if err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	if err := create(ctx, tx, transaction); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}
	go metric.Done(false)

	return nil
}

func (c *Client) ReadTransaction(ctx context.Context, id int64) (*models.Transaction, db.Error) {
	metric := c.metrics.NewDBQuery("ReadTransaction")

	transaction := new(models.Transaction)
	err := newTransactionQ(c.bun, transaction).Where("id = ?", id).Scan(ctx)
	if err != nil {
		dbErr := c.bun.ProcessError(err)

		if errors.Is(dbErr, db.ErrNoEntries) {
			// report no entries as a non error
			go metric.Done(false)
		} else {
			go metric.Done(true)
		}

		return nil, dbErr
	}
	go metric.Done(false)

	return transaction, nil
}

func (c *Client) ReadTransactionsPage(ctx context.Context, index, count int) ([]*models.Transaction, db.Error) {
	metric := c.metrics.NewDBQuery("ReadTransactionsPage")

	var transaction []*models.Transaction
	err := newTransactionsQ(c.bun, &transaction).
		Limit(count).
		Offset(libdatabase.Offset(index, count)).
		Scan(ctx)
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}
	go metric.Done(false)

	return transaction, nil
}

func newTransactionQ(c bun.IDB, transaction *models.Transaction) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(transaction)
}

func newTransactionsQ(c bun.IDB, transaction *[]*models.Transaction) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(transaction)
}
