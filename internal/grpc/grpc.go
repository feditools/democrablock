package grpc

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/models"
	libgrpc "github.com/feditools/go-lib/grpc"
	loginGrpc "github.com/feditools/login/pkg/grpc"
	"github.com/spf13/viper"
	"time"
)

const (
	fediAccountLifeWindow         = 10 * time.Minute
	fediAccountCleanWindow        = 5 * time.Minute
	fediAccountMaxEntriesInWindow = 10000

	fediInstanceLifeWindow         = 10 * time.Minute
	fediInstanceCleanWindow        = 5 * time.Minute
	fediInstanceMaxEntriesInWindow = 10000
)

type Client struct {
	login *loginGrpc.Client

	fediAccount   *bigcache.BigCache
	fediInstances *bigcache.BigCache
}

func New(ctx context.Context) (*Client, error) {
	cred := libgrpc.NewCredential(viper.GetString(config.Keys.GRPCLoginToken))
	loginClient, err := loginGrpc.NewClient(viper.GetString(config.Keys.GRPCLoginAddress), cred)
	if err != nil {
		return nil, err
	}
	loginResp, err := loginClient.Ping(ctx)
	if err != nil {
		return nil, err
	}
	if loginResp.Response != "pong" {
		return nil, fmt.Errorf("login returns invalid response to ping: %s", loginResp.Response)
	}

	// make caches
	gob.Register(models.FediAccount{})
	gob.Register(models.FediInstance{})

	fediAccount, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32, // nolint
		LifeWindow:         fediAccountLifeWindow,
		CleanWindow:        fediAccountCleanWindow,
		MaxEntriesInWindow: fediAccountMaxEntriesInWindow,
		MaxEntrySize:       500, // nolint
		Verbose:            true,
		HardMaxCacheSize:   8192, // nolint
	})
	if err != nil {
		return nil, err
	}

	fediInstances, err := bigcache.NewBigCache(bigcache.Config{
		Shards:             32, // nolint
		LifeWindow:         fediInstanceLifeWindow,
		CleanWindow:        fediInstanceCleanWindow,
		MaxEntriesInWindow: fediInstanceMaxEntriesInWindow,
		MaxEntrySize:       500, // nolint
		Verbose:            true,
		HardMaxCacheSize:   8192, // nolint
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		login: loginClient,

		fediAccount:   fediAccount,
		fediInstances: fediInstances,
	}, nil
}
