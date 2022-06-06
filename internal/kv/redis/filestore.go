package redis

import (
	"context"
	"time"

	"github.com/feditools/democrablock/internal/kv"
)

// filestore presigned url tokens

// DeleteFileStorePresignedURL deletes a filestore presigned url token from redis.
func (c *Client) DeleteFileStorePresignedURL(ctx context.Context, token string) kv.Error {
	_, err := c.redis.Del(ctx, kv.KeyFileStorePresignedURL(token)).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}

// GetFileStorePresignedURL retrieves a filestore presigned url token from redis.
func (c *Client) GetFileStorePresignedURL(ctx context.Context, token string) (string, kv.Error) {
	resp, err := c.redis.Get(ctx, kv.KeyFileStorePresignedURL(token)).Result()
	if err != nil {
		return "", c.ProcessError(err)
	}

	return resp, nil
}

// SetFileStorePresignedURL adds a filestore presigned url token to redis.
func (c *Client) SetFileStorePresignedURL(ctx context.Context, token string, url string, expire time.Duration) kv.Error {
	_, err := c.redis.SetEX(ctx, kv.KeyFileStorePresignedURL(token), url, expire).Result()
	if err != nil {
		return c.ProcessError(err)
	}

	return nil
}
