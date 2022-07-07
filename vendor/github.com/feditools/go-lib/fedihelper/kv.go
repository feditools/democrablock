package fedihelper

import (
	"context"
	"time"
)

type KV interface {
	// access token

	DeleteAccessToken(ctx context.Context, accountID int64) (err error)
	GetAccessToken(ctx context.Context, accountID int64) (accessToken string, err error)
	SetAccessToken(ctx context.Context, accountID int64, accessToken string) (err error)

	// instance oauth

	DeleteInstanceOAuth(ctx context.Context, instanceID int64) (err error)
	GetInstanceOAuth(ctx context.Context, instanceID int64) (clientID string, clientSecret string, err error)
	SetInstanceOAuth(ctx context.Context, instanceID int64, clientID string, clientSecret string) (err error)

	// federated instance node info

	DeleteFediNodeInfo(ctx context.Context, domain string) (err error)
	GetFediNodeInfo(ctx context.Context, domain string) (nodeinfo string, err error)
	SetFediNodeInfo(ctx context.Context, domain string, nodeinfo string, expire time.Duration) (err error)
}
