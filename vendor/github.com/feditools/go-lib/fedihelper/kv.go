package fedihelper

import (
	"context"
	"time"
)

type KV interface {
	// federated instance node info

	DeleteFediNodeInfo(ctx context.Context, domain string) (err error)
	GetFediNodeInfo(ctx context.Context, domain string) (nodeinfo string, err error)
	SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) (err error)
}
