package cachemem

import (
	"context"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/models"
)

func (c *CacheMem) CountTransactions(ctx context.Context) (count int64, err db.Error) {
	return c.db.CountTransactions(ctx)
}

func (c *CacheMem) CreateTransaction(ctx context.Context, transaction *models.Transaction) (err db.Error) {
	return c.db.CreateTransaction(ctx, transaction)
}

func (c *CacheMem) CreateTransactionTX(ctx context.Context, txID db.TxID, transaction *models.Transaction) (err db.Error) {
	return c.db.CreateTransactionTX(ctx, txID, transaction)
}

func (c *CacheMem) ReadTransaction(ctx context.Context, id int64) (transaction *models.Transaction, err db.Error) {
	return c.db.ReadTransaction(ctx, id)
}

func (c *CacheMem) ReadTransactionsPage(ctx context.Context, index, count int) (transactions []*models.Transaction, err db.Error) {
	return c.db.ReadTransactionsPage(ctx, index, count)
}
