package immudb

import (
	"context"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/models"
)

func (c *Client) CountTransactions(ctx context.Context) (count int64, err db.Error) {
	// TODO implement me
	panic("implement me")
}

func (c *Client) CreateTransaction(ctx context.Context, transaction *models.Transaction) (err db.Error) {
	// TODO implement me
	panic("implement me")
}

func (c *Client) CreateTransactionTX(ctx context.Context, txID db.TxID, transaction *models.Transaction) (err db.Error) {
	// TODO implement me
	panic("implement me")
}

func (c *Client) ReadTransaction(ctx context.Context, id int64) (transaction *models.Transaction, err db.Error) {
	// TODO implement me
	panic("implement me")
}

func (c *Client) ReadTransactionsPage(ctx context.Context, index, count int) (transactions []*models.Transaction, err db.Error) {
	// TODO implement me
	panic("implement me")
}
