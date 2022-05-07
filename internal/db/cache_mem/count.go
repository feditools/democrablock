package cachemem

import (
	"context"
	"encoding/binary"
	"errors"

	"github.com/allegro/bigcache/v3"
)

const (
	tableNameBlocks = "blocks"
)

func keyCountBlocks() string {
	return tableNameBlocks
}

func (c *CacheMem) getCount(_ context.Context, k string) (int64, bool) {
	l := logger.WithField("func", "getCount")

	// check domain cache
	entry, err := c.count.Get(k)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return 0, false
	}
	if err != nil {
		l.Warnf("cache get: %s", err.Error())

		return 0, false
	}
	i := int64(binary.LittleEndian.Uint64(entry))

	return i, true
}

func (c *CacheMem) setCount(_ context.Context, k string, count int64) {
	l := logger.WithField("func", "setCount")

	// encode count
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(count))
	if err := c.count.Set(k, b); err != nil {
		l.Warnf("cache set: %s", err.Error())

		return
	}
}
