package immudb

import (
	"bytes"
	"context"
	"time"

	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/feditools/democrablock/internal/util"

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
		l.Errorf("NewTx: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
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
		l.Errorf("Commit: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = createdAt
	account.ID = resp.GetLastInsertedPKs()[TableNameFediAccounts].GetN()

	go metric.Done(false)

	return nil
}

func (c *Client) IncFediAccountLoginCount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("IncFediAccountLoginCount")
	l := logger.WithField("func", "IncFediAccountLoginCount")

	entry, err := c.db.Get(ctx, KeyFediAccountLoginCount(account.ID))
	if err != nil && err.Error() != KeyNotFoundError {
		l.Errorf("Get: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
	}

	var preconditions []*schema.Precondition
	count := int64(1)
	now := time.Now().UTC()
	if entry != nil {
		precondition := schema.PreconditionKeyNotModifiedAfterTX(
			KeyFediAccountLoginCount(account.ID),
			entry.Tx,
		)
		count = util.BytesToInt64(entry.Value) + 1

		preconditions = append(preconditions, precondition)
	}
	_, err = c.db.SetAll(ctx, &schema.SetRequest{
		KVs: []*schema.KeyValue{
			{
				Key:   KeyFediAccountLoginCount(account.ID),
				Value: util.Int64ToBytes(count),
			},
			{
				Key:   KeyFediAccountLoginLast(account.ID),
				Value: util.TimeToBytes(now),
			},
		},
		Preconditions: preconditions,
	})
	if err != nil {
		l.Errorf("SetAll: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
	}

	account.LogInCount = count
	account.LogInLast = now

	go metric.Done(false)

	return nil
}

func (c *Client) ReadFediAccount(ctx context.Context, id int64) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccount")
	l := logger.WithField("func", "ReadFediAccount")

	resp, err := c.db.SQLQuery(ctx, selectFediAccount(id), nil, true)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	if len(resp.GetRows()) == 0 {
		go metric.Done(false)

		return nil, db.ErrNoEntries
	}

	// make new account from
	account := makeFediAccountFromRow(resp.GetRows()[0])

	// get login info
	loginCount, loginLast, err := c.readFediAccountLoginInfo(ctx, account.ID)
	if err != nil {
		l.Errorf("read login info: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}
	if loginCount > 0 {
		account.LogInCount = loginCount
		account.LogInLast = loginLast
	}

	go metric.Done(false)

	return account, nil
}

func (c *Client) ReadFediAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccountByUsername")
	l := logger.WithField("func", "ReadFediAccountByUsername")

	resp, err := c.db.SQLQuery(ctx, selectFediAccountByUsername(instanceID, username), nil, true)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	if len(resp.GetRows()) == 0 {
		go metric.Done(false)

		return nil, db.ErrNoEntries
	}

	// make new account from
	account := makeFediAccountFromRow(resp.GetRows()[0])

	// get login info
	loginCount, loginLast, err := c.readFediAccountLoginInfo(ctx, account.ID)
	if err != nil {
		l.Errorf("read login info: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}
	if loginCount > 0 {
		account.LogInCount = loginCount
		account.LogInLast = loginLast
	}

	go metric.Done(false)

	return account, nil
}

func (c *Client) ReadFediAccountsPage(ctx context.Context, index, count int) ([]*models.FediAccount, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediAccountsPage")
	l := logger.WithField("func", "ReadFediAccountsPage")

	lastReadID, err := c.PageHelper(ctx, TableNameFediAccounts, index, count)
	if err != nil {
		l.Errorf("page helper: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	l.Debugf("last seen id: %d", lastReadID)

	resp, err := c.db.SQLQuery(
		ctx,
		selectFediAccountsPage(lastReadID, count, "id", true),
		nil,
		true,
	)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	accounts := make([]*models.FediAccount, len(resp.GetRows()))
	for i, row := range resp.GetRows() {
		// make new account from
		account := makeFediAccountFromRow(row)

		// get login info
		loginCount, loginLast, err := c.readFediAccountLoginInfo(ctx, account.ID)
		if err != nil {
			l.Errorf("read login info: %s", err.Error())
			go metric.Done(true)

			return nil, c.ProcessError(err)
		}
		if loginCount > 0 {
			account.LogInCount = loginCount
			account.LogInLast = loginLast
		}

		accounts[i] = account
	}

	go metric.Done(false)

	return accounts, nil
}

func (c *Client) UpdateFediAccount(ctx context.Context, account *models.FediAccount) db.Error {
	metric := c.metrics.NewDBQuery("UpdateFediAccount")

	go metric.Done(false)

	return nil
}

// privates

func makeFediAccountFromRow(row *schema.Row) *models.FediAccount {
	newAccount := models.FediAccount{
		ID:          row.GetValues()[fediAccountColumnID].GetN(),
		CreatedAt:   tsToTime(row.GetValues()[fediAccountColumnCreatedAt].GetTs()),
		UpdatedAt:   tsToTime(row.GetValues()[fediAccountColumnUpdatedAt].GetTs()),
		Username:    row.GetValues()[fediAccountColumnUsername].GetS(),
		InstanceID:  row.GetValues()[fediAccountColumnInstanceID].GetN(),
		ActorURI:    row.GetValues()[fediAccountColumnActorURI].GetS(),
		DisplayName: row.GetValues()[fediAccountColumnDisplayName].GetS(),
		LastFinger:  tsToTime(row.GetValues()[fediAccountColumnLastFinger].GetTs()),
		Admin:       row.GetValues()[fediAccountColumnIsAdmin].GetB(),
	}
	if !isNull(row.GetValues()[fediAccountColumnAccessToken]) {
		newAccount.AccessToken = row.GetValues()[fediAccountColumnAccessToken].GetBs()
	}

	return &newAccount
}

func (c *Client) readFediAccountLoginInfo(ctx context.Context, id int64) (int64, time.Time, error) {
	l := logger.WithField("func", "readFediAccountLoginInfo")

	entries, err := c.db.GetAll(
		ctx,
		[][]byte{
			KeyFediAccountLoginCount(id),
			KeyFediAccountLoginLast(id),
		},
	)
	if err != nil {
		l.Errorf("get: %s", err.Error())

		return 0, time.Time{}, err
	}

	count := int64(0)
	last := time.Time{}
	for _, entry := range entries.GetEntries() {
		switch {
		case bytes.Equal(entry.GetKey(), KeyFediAccountLoginCount(id)):
			count = util.BytesToInt64(entry.GetValue())
		case bytes.Equal(entry.GetKey(), KeyFediAccountLoginLast(id)):
			last = util.BytesToTime(entry.GetValue())
		}
	}

	return count, last, nil
}
