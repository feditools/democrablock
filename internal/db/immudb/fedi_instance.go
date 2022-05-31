package immudb

import (
	"context"

	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/models"
)

func (c *Client) CountFediInstances(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountFediInstances")

	go metric.Done(false)

	return 0, nil
}

func (c *Client) CreateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	metric := c.metrics.NewDBQuery("CreateFediInstance")

	go metric.Done(false)

	return nil
}

func (c *Client) ReadFediInstance(ctx context.Context, id int64) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstance")

	go metric.Done(false)

	return nil, nil
}

func (c *Client) ReadFediInstanceByDomain(ctx context.Context, domain string) (*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstanceByDomain")

	go metric.Done(false)

	return nil, nil
}

func (c *Client) ReadFediInstancesPage(ctx context.Context, index, count int) ([]*models.FediInstance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadFediInstancesPage")

	go metric.Done(false)

	return nil, nil
}

func (c *Client) UpdateFediInstance(ctx context.Context, instance *models.FediInstance) db.Error {
	metric := c.metrics.NewDBQuery("UpdateFediInstance")

	go metric.Done(false)

	return nil
}
