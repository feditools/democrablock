package filestore

import (
	"context"
	"net/url"
)

type FileStore interface {
	GetEvidenceFile(ctx context.Context, hash []byte, kind string) ([]byte, error)
	GetEvidencePresignedURL(ctx context.Context, hash []byte, kind string) (*url.URL, error)
	PutEvidenceFile(ctx context.Context, hash []byte, kind string, data []byte) error
}
