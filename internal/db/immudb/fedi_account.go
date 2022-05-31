package immudb

import (
	"context"
	"time"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/models"
)

func (c *Client) CountFediAccounts(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountFediAccounts")
	l := logger.WithField("func", "CountFediAccounts")

	resp, err := c.db.SQLQuery(ctx, countFediAccounts(), nil, true)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return 0, c.ProcessError(err)
	}

	go metric.Done(false)

	return resp.GetRows()[0].GetValues()[0].GetN(), nil
}

func (c *Client) CountFediAccountsForInstance(ctx context.Context, instanceID int64) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountFediAccountsForInstance")
	l := logger.WithField("func", "CountFediAccountsForInstance")

	resp, err := c.db.SQLQuery(ctx, countFediAccountsForInstance(instanceID), nil, true)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return 0, c.ProcessError(err)
	}

	go metric.Done(false)

	return resp.GetRows()[0].GetValues()[0].GetN(), nil
}

func (c *Client) CreateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("CreateFediAccount")
	l := logger.WithField("func", "CreateFediAccount")

	createdAt := time.Now().UTC()

	// create transaction
	tx, err := c.db.NewTx(ctx)
	if err != nil {
		return err
	}

	err = tx.SQLExec(ctx, insertFediAccount(account, createdAt), nil)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
	}

	// commit
	resp, err := tx.Commit(ctx)
	if err != nil {
		return err
	}

	l.Debugf("inserted pk; %#v", resp.GetLastInsertedPKs())

	account.CreatedAt = createdAt
	account.UpdatedAt = createdAt
	account.ID = resp.GetLastInsertedPKs()[TableNameFediAccounts].GetN()

	go metric.Done(false)

	return nil
}

func (c *Client) IncFediAccountLoginCount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("IncFediAccountLoginCount")

	go metric.Done(false)

	return nil
}

func (c *Client) ReadFediAccount(ctx context.Context, id int64) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccount")

	go metric.Done(false)

	return nil, nil
}

func (c *Client) ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccountByUsername")

	go metric.Done(false)

	return nil, nil
}

func (c *Client) ReadFediAccountsPage(ctx context.Context, index, count int) ([]*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccountsPage")

	go metric.Done(false)

	return nil, nil
}

func (c *Client) UpdateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("UpdateFediAccount")

	go metric.Done(false)

	return nil
}
