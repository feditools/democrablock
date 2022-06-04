package minio

import (
	"context"
	"net/url"
)

func (m Module) GetFile(ctx context.Context, group string, hash []byte, suffix string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (m Module) GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, error) {
	// TODO implement me
	panic("implement me")
}

func (m Module) PutFile(ctx context.Context, group string, hash []byte, suffix string, data []byte) error {
	// TODO implement me
	panic("implement me")
}
