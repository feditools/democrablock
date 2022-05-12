package grpc

import (
	"context"
	"github.com/feditools/democrablock/internal/models"
)

func (c *Client) GetFediAccount(ctx context.Context, id int64) (*models.FediAccount, error) {
	gFediAccount, err := c.login.GetFediAccount(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.FediAccount{
		ID:          gFediAccount.Id,
		Username:    gFediAccount.Username,
		InstanceID:  gFediAccount.InstanceId,
		DisplayName: gFediAccount.DisplayName,
		Admin:       gFediAccount.IsAdmin,
	}, nil
}
