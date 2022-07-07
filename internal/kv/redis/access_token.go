package redis

import (
	"context"
	"github.com/feditools/democrablock/internal/kv"
	"github.com/feditools/democrablock/internal/util"
)

func (c *Client) DeleteAccessToken(ctx context.Context, userID int64) error {
	_, err := c.redis.Del(ctx, kv.KeyUserAccessToken(userID)).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}

func (c *Client) GetAccessToken(ctx context.Context, userID int64) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyUserAccessToken(userID)).Bytes()
	if err != nil {
		return "", c.ProcessError(err)
	}

	data, err := util.Decrypt(resp)
	if err != nil {
		return "", kv.NewEncryptionError(err.Error())
	}

	return string(data), nil
}

func (c *Client) SetAccessToken(ctx context.Context, userID int64, accessToken string) error {
	data, err := util.Encrypt([]byte(accessToken))
	if err != nil {
		return kv.NewEncryptionError(err.Error())
	}

	_, err = c.redis.Set(ctx, kv.KeyUserAccessToken(userID), data, 0).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}
