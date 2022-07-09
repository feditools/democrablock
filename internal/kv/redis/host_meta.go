package redis

import (
	"context"
	"time"

	"github.com/feditools/democrablock/internal/kv"
)

// DeleteHostMeta deletes fedi host meta from redis.
func (c *Client) DeleteHostMeta(ctx context.Context, domain string) error {
	_, err := c.redis.Del(ctx, kv.KeyFediNodeInfo(domain)).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}

// GetHostMeta retrieves fedi host meta from redis.
func (c *Client) GetHostMeta(ctx context.Context, domain string) ([]byte, error) {
	resp, err := c.redis.Get(ctx, kv.KeyFediNodeInfo(domain)).Bytes()
	if err != nil {
		return nil, c.ProcessError(err)
	}

	return resp, nil
}

// SetHostMeta adds fedi host meta to redis.
func (c *Client) SetHostMeta(ctx context.Context, domain string, hostmeta []byte, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyFediNodeInfo(domain), hostmeta, expire).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}
