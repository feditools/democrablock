package fedihelper

import (
	"context"
	"time"
)

type KV interface {
	// access token

	DeleteAccessToken(ctx context.Context, userID int64) (err error)
	GetAccessToken(ctx context.Context, userID int64) (accessToken string, err error)
	SetAccessToken(ctx context.Context, userID int64, accessToken string) (err error)

	// federated instance node info

	DeleteFediNodeInfo(ctx context.Context, domain string) (err error)
	GetFediNodeInfo(ctx context.Context, domain string) (nodeinfo string, err error)
	SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) (err error)
}
