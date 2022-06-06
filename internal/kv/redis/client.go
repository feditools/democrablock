package redis

import (
	"context"

	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/kv"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/spf13/viper"
)

// New creates a new redis client.
func New(ctx context.Context) (*Client, error) {
	l := logger.WithField("func", "New")

	r := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(config.Keys.RedisAddress),
		Password: viper.GetString(config.Keys.RedisPassword),
		DB:       viper.GetInt(config.Keys.RedisDB),
	})

	pool := goredis.NewPool(r)
	c := Client{
		redis: r,
		sync:  redsync.New(pool),
	}

	resp := c.redis.Ping(ctx)
	l.Debugf("%s", resp.String())

	return &c, nil
}

// Client represents a redis client.
type Client struct {
	redis *redis.Client
	sync  *redsync.Redsync
}

// Close closes the redis pool.
func (c *Client) Close(_ context.Context) kv.Error {
	return c.redis.Close()
}

// RedisClient returns the redis client.
func (c *Client) RedisClient() *redis.Client {
	return c.redis
}
