package kv

import (
	"context"

	"github.com/feditools/go-lib/fedihelper"
)

// KV represents a key value store.
type KV interface {
	fedihelper.KV

	Close(ctx context.Context) error
}
