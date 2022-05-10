package cachemem

import (
	"context"
	"time"

	bigcache "github.com/allegro/bigcache/v3"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/metrics"
)

// CacheMem is an in memory caching middleware for our db interface.
type CacheMem struct {
	db      db.DB
	metrics metrics.Collector

	count *bigcache.BigCache

	allCaches []*bigcache.BigCache
}

// New creates a new in memory cache.
func New(_ context.Context, d db.DB, m metrics.Collector) (db.DB, error) {
	l := logger.WithField("func", "New")

	//revive:disable:add-constant
	count, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         30 * time.Second,
		CleanWindow:        1 * time.Minute,
		MaxEntriesInWindow: 10000,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		l.Errorf("count: %s", err.Error())

		return nil, err
	}
	//revive:enable:add-constant

	return &CacheMem{
		db:      d,
		metrics: m,

		count: count,

		allCaches: []*bigcache.BigCache{
			count,
		},
	}, nil
}

// Close is a pass through.
func (c *CacheMem) Close(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Close()
	}

	return c.db.Close(ctx)
}

// Create is a pass through.
func (c *CacheMem) Create(ctx context.Context, i interface{}) db.Error {
	return c.db.Create(ctx, i)
}

// DoMigration is a pass through.
func (c *CacheMem) DoMigration(ctx context.Context) db.Error {
	return c.db.DoMigration(ctx)
}

// LoadTestData is a pass through.
func (c *CacheMem) LoadTestData(ctx context.Context) db.Error {
	return c.db.LoadTestData(ctx)
}

// ReadByID is a pass through.
func (c *CacheMem) ReadByID(ctx context.Context, id int64, i interface{}) db.Error {
	return c.db.ReadByID(ctx, id, i)
}

// ResetCache clears all the caches.
func (c *CacheMem) ResetCache(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Reset()
	}

	return c.db.ResetCache(ctx)
}

// Update is a pass through.
func (c *CacheMem) Update(ctx context.Context, i interface{}) db.Error {
	return c.db.Update(ctx, i)
}
