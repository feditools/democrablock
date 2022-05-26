package redis

import (
	"context"
	"github.com/feditools/democrablock/internal/kv"
	"time"
)

// fedi node info

// DeleteFediNodeInfo deletes fedi nodeinfo from redis.
func (c *Client) DeleteFediNodeInfo(ctx context.Context, domain string) error {
	_, err := c.redis.Del(ctx, kv.KeyFediNodeInfo(domain)).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetFediNodeInfo retrieves fedi nodeinfo from redis.
func (c *Client) GetFediNodeInfo(ctx context.Context, domain string) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyFediNodeInfo(domain)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// SetFediNodeInfo adds fedi nodeinfo to redis.
func (c *Client) SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyFediNodeInfo(domain), nodeinfo, expire).Result()
	if err != nil {
		return err
	}

	return nil
}
