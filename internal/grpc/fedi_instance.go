package grpc

import (
	"context"
	"github.com/feditools/democrablock/internal/models"
)

func (c *Client) GetFediInstance(ctx context.Context, id int64) (*models.FediInstance, error) {
	gFediInstance, err := c.login.GetFediInstance(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.FediInstance{
		ID:             gFediInstance.Id,
		Domain:         gFediInstance.Domain,
		ServerHostname: gFediInstance.ServerHostname,
		Software:       gFediInstance.Software,
	}, nil
}
