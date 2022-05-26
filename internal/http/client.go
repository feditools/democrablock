package http

import (
	"context"
	"fmt"
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/viper"
)

func NewClient(_ context.Context) (*Client, error) {
	userAgent := fmt.Sprintf("Go-http-client/2.0 (%s/%s; +%s/)",
		viper.GetString(config.Keys.ApplicationName),
		viper.GetString(config.Keys.SoftwareVersion),
		viper.GetString(config.Keys.ServerExternalURL),
	)

	return &Client{
		userAgent: userAgent,
	}, nil
}

type Client struct {
	userAgent string
}
