package immudb

import (
	"context"
	"time"

	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/feditools/democrablock/internal/db/immudb/statements"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/models"
)

func (c *Client) CountFediInstances(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountFediInstances")
	l := logger.WithField("func", "CountFediInstances")

	// run query
	resp, err := c.db.SQLQuery(
		ctx,
		statements.CountFediInstances(),
		nil,
		true,
	)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return 0, c.ProcessError(err)
	}

	go metric.Done(false)

	return resp.GetRows()[0].GetValues()[0].GetN(), nil
}

func (c *Client) CreateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	metric := c.metrics.NewDBQuery("CreateFediInstance")
	l := logger.WithField("func", "CreateFediInstance")

	// prep params
	createdAt := time.Now().UTC()
	params := map[string]interface{}{
		statements.FediInstanceColumnNameCreatedAt:      createdAt,
		statements.FediInstanceColumnNameUpdatedAt:      createdAt,
		statements.FediInstanceColumnNameDomain:         instance.Domain,
		statements.FediInstanceColumnNameActorURI:       instance.ActorURI,
		statements.FediInstanceColumnNameServerHostname: instance.ServerHostname,
		statements.FediInstanceColumnNameSoftware:       instance.Software,
	}
	if instance.ClientID != "" {
		params[statements.FediInstanceColumnNameClientID] = instance.ClientID
	} else {
		params[statements.FediInstanceColumnNameClientID] = nil
	}
	if len(instance.ClientSecret) > 0 {
		params[statements.FediInstanceColumnNameClientSecret] = instance.ClientSecret
	} else {
		params[statements.FediInstanceColumnNameClientSecret] = nil
	}

	// run query
	resp, err := c.db.SQLExec(
		ctx,
		statements.InsertFediInstance(),
		params,
	)
	if err != nil {
		l.Errorf("Commit: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
	}

	instance.CreatedAt = createdAt
	instance.UpdatedAt = createdAt
	instance.ID = resp.LastInsertedPk()[statements.FediInstancesTableName].GetN()

	go metric.Done(false)

	return nil
}

func (c *Client) ReadFediInstance(ctx context.Context, id int64) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstance")
	l := logger.WithField("func", "ReadFediInstance")

	// prep params
	params := map[string]interface{}{
		statements.FediInstanceColumnNameID: id,
	}

	// run query
	resp, err := c.db.SQLQuery(
		ctx,
		statements.SelectFediInstance(),
		params,
		true,
	)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	if len(resp.GetRows()) == 0 {
		go metric.Done(false)

		return nil, db.ErrNoEntries
	}

	// make new instance from
	instance := makeFediInstanceFromRow(resp.GetRows()[0])

	go metric.Done(false)

	return instance, nil
}

func (c *Client) ReadFediInstanceByDomain(ctx context.Context, domain string) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstanceByDomain")
	l := logger.WithField("func", "ReadFediInstanceByDomain")

	// prep params
	params := map[string]interface{}{
		statements.FediInstanceColumnNameDomain: domain,
	}

	// run query
	resp, err := c.db.SQLQuery(
		ctx,
		statements.SelectFediInstanceByDomain(),
		params,
		true,
	)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	if len(resp.GetRows()) == 0 {
		go metric.Done(false)

		return nil, db.ErrNoEntries
	}

	// make new instance from
	instance := makeFediInstanceFromRow(resp.GetRows()[0])

	go metric.Done(false)

	return instance, nil
}

func (c *Client) ReadFediInstancesPage(ctx context.Context, index, count int) ([]*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstancesPage")
	l := logger.WithField("func", "ReadFediInstancesPage")

	lastReadID, err := c.PageHelper(ctx, statements.FediInstancesTableName, index, count)
	if err != nil {
		l.Errorf("page helper: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	l.Debugf("last seen id: %d", lastReadID)

	// prep params
	params := map[string]interface{}{
		statements.ParamLastReadID: lastReadID,
	}

	// run query
	resp, err := c.db.SQLQuery(
		ctx,
		statements.SelectFediInstancesPage(true, count),
		params,
		true,
	)
	if err != nil {
		l.Errorf("SQLQuery: %s", err.Error())
		go metric.Done(true)

		return nil, c.ProcessError(err)
	}

	accounts := make([]*models.FediInstance, len(resp.GetRows()))
	for i, row := range resp.GetRows() {
		accounts[i] = makeFediInstanceFromRow(row)
	}

	go metric.Done(false)

	return accounts, nil
}

func (c *Client) UpdateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	metric := c.metrics.NewDBQuery("UpdateFediInstance")
	l := logger.WithField("func", "UpdateFediInstance")

	// prep params
	updatedAt := time.Now().UTC()
	params := map[string]interface{}{
		statements.FediInstanceColumnNameID:             instance.ID,
		statements.FediInstanceColumnNameCreatedAt:      instance.CreatedAt,
		statements.FediInstanceColumnNameUpdatedAt:      updatedAt,
		statements.FediInstanceColumnNameDomain:         instance.Domain,
		statements.FediInstanceColumnNameActorURI:       instance.ActorURI,
		statements.FediInstanceColumnNameServerHostname: instance.ServerHostname,
		statements.FediInstanceColumnNameSoftware:       instance.Software,
	}
	if instance.ClientID != "" {
		params[statements.FediInstanceColumnNameClientID] = instance.ClientID
	} else {
		params[statements.FediInstanceColumnNameClientID] = nil
	}
	if len(instance.ClientSecret) > 0 {
		params[statements.FediInstanceColumnNameClientSecret] = instance.ClientSecret
	} else {
		params[statements.FediInstanceColumnNameClientSecret] = nil
	}

	// run query
	l.Debugf("statement:%s\nparams:\n%+v", statements.UpsertFediAccount(), params)
	_, err := c.db.SQLExec(
		ctx,
		statements.UpsertFediInstance(),
		params,
	)
	if err != nil {
		l.Errorf("SQLExec: %s", err.Error())
		go metric.Done(true)

		return c.ProcessError(err)
	}

	instance.UpdatedAt = updatedAt

	go metric.Done(false)

	return nil
}

// privates

func makeFediInstanceFromRow(row *schema.Row) *models.FediInstance {
	newInstance := models.FediInstance{
		ID:             row.GetValues()[statements.FediInstanceColumnIndexID].GetN(),
		CreatedAt:      tsToTime(row.GetValues()[statements.FediInstanceColumnIndexCreatedAt].GetTs()),
		UpdatedAt:      tsToTime(row.GetValues()[statements.FediInstanceColumnIndexUpdatedAt].GetTs()),
		Domain:         row.GetValues()[statements.FediInstanceColumnIndexDomain].GetS(),
		ActorURI:       row.GetValues()[statements.FediInstanceColumnIndexActorURI].GetS(),
		ServerHostname: row.GetValues()[statements.FediInstanceColumnIndexServerHostname].GetS(),
		Software:       row.GetValues()[statements.FediInstanceColumnIndexSoftware].GetS(),
	}
	if !isNull(row.GetValues()[statements.FediInstanceColumnIndexClientID]) {
		newInstance.ClientID = row.GetValues()[statements.FediInstanceColumnIndexClientID].GetS()
	}
	if !isNull(row.GetValues()[statements.FediInstanceColumnIndexClientSecret]) {
		newInstance.ClientSecret = row.GetValues()[statements.FediInstanceColumnIndexClientSecret].GetBs()
	}

	return &newInstance
}
