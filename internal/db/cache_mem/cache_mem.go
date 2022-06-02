package cachemem

import (
	"context"
	"time"

	"github.com/feditools/go-lib/metrics"

	bigcache "github.com/allegro/bigcache/v3"
	"github.com/feditools/democrablock/internal/db"
)

const (
	fediAccountLifeWindow         = 10 * time.Minute
	fediAccountCleanWindow        = 5 * time.Minute
	fediAccountMaxEntriesInWindow = 10000

	fediInstanceLifeWindow         = 10 * time.Minute
	fediInstanceCleanWindow        = 5 * time.Minute
	fediInstanceMaxEntriesInWindow = 10000
)

// CacheMem is an in memory caching middleware for our db interface.
type CacheMem struct {
	db      db.DB
	metrics metrics.Collector

	count *bigcache.BigCache

	fediAccount             *bigcache.BigCache
	fediAccountUsernameToID *bigcache.BigCache

	fediInstance           *bigcache.BigCache
	fediInstanceDomainToID *bigcache.BigCache

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

	fediAccount, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediAccountLifeWindow,
		CleanWindow:        fediAccountCleanWindow,
		MaxEntriesInWindow: fediAccountMaxEntriesInWindow,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediAccountUsernameToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediAccountLifeWindow,
		CleanWindow:        fediAccountCleanWindow,
		MaxEntriesInWindow: fediAccountMaxEntriesInWindow,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediInstance, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediInstanceLifeWindow,
		CleanWindow:        fediInstanceCleanWindow,
		MaxEntriesInWindow: fediInstanceMaxEntriesInWindow,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}

	fediInstanceDomainToID, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32,
		LifeWindow:         fediInstanceLifeWindow,
		CleanWindow:        fediInstanceCleanWindow,
		MaxEntriesInWindow: fediInstanceMaxEntriesInWindow,
		MaxEntrySize:       8,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	})
	if err != nil {
		return nil, err
	}
	//revive:enable:add-constant

	return &CacheMem{
		db:      d,
		metrics: m,

		count: count,

		fediAccount:             fediAccount,
		fediAccountUsernameToID: fediAccountUsernameToID,

		fediInstance:           fediInstance,
		fediInstanceDomainToID: fediInstanceDomainToID,

		allCaches: []*bigcache.BigCache{
			count,
			fediAccount,
			fediAccountUsernameToID,
			fediInstance,
			fediInstanceDomainToID,
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

// DoMigration is a pass through.
func (c *CacheMem) DoMigration(ctx context.Context) db.Error {
	return c.db.DoMigration(ctx)
}

// LoadTestData is a pass through.
func (c *CacheMem) LoadTestData(ctx context.Context) db.Error {
	return c.db.LoadTestData(ctx)
}

// ResetCache clears all the caches.
func (c *CacheMem) ResetCache(ctx context.Context) db.Error {
	for _, cache := range c.allCaches {
		_ = cache.Reset()
	}

	return c.db.ResetCache(ctx)
}
