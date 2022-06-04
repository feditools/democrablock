package redis

import (
	"context"
	"github.com/feditools/democrablock/internal/kv"
	"time"
)

// filestore presigned url tokens

// DeleteFileStorePresignedURL deletes a filestore presigned url token from redis.
func (c *Client) DeleteFileStorePresignedURL(ctx context.Context, token string) error {
	_, err := c.redis.Del(ctx, kv.KeyFileStorePresignedURL(token)).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetFileStorePresignedURL retrieves a filestore presigned url token from redis.
func (c *Client) GetFileStorePresignedURL(ctx context.Context, token string) (string, error) {
	resp, err := c.redis.Get(ctx, kv.KeyFileStorePresignedURL(token)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// SetFileStorePresignedURL adds a filestore presigned url token to redis.
func (c *Client) SetFileStorePresignedURL(ctx context.Context, token string, url string, expire time.Duration) error {
	_, err := c.redis.SetEX(ctx, kv.KeyFileStorePresignedURL(token), url, expire).Result()
	if err != nil {
		return err
	}

	return nil
}
