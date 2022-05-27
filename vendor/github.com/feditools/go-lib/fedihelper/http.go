package fedihelper

import (
	"context"
	"net/http"
)

type HTTP interface {
	Get(ctx context.Context, url string) (resp *http.Response, err error)
	Transport() (transport http.RoundTripper)
}
