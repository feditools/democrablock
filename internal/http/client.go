package http

import (
	"context"
	"fmt"
	"io"
	"net/http"

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

// Get calls http.Get with expected http User-Agent.
func (c *Client) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	client := &http.Client{}
	req, err := c.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

// NewRequest calls http.NewRequest with expected http User-Agent.
func (c *Client) NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}

// Transport adds the expected http User-Agent to any request.
type Transport struct {
	userAgent string
}

func (c *Client) Transport() (transport http.RoundTripper) {
	return &Transport{userAgent: c.userAgent}
}

// RoundTrip executes the default http.Transport with expected http User-Agent.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)

	return http.DefaultTransport.RoundTrip(req)
}
