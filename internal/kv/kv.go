package kv

import (
	"context"
	"time"

	"github.com/feditools/go-lib/fedihelper"
)

// KV represents a key value store.
type KV interface {
	fedihelper.KV

	Close(ctx context.Context) error

	// filestore presigned url tokens

	DeleteFileStorePresignedURL(ctx context.Context, token string) (err error)
	GetFileStorePresignedURL(ctx context.Context, token string) (url string, err error)
	SetFileStorePresignedURL(ctx context.Context, token string, url string, expire time.Duration) (err error)
}
