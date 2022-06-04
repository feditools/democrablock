package filestore

import (
	"context"
	"fmt"
	"net/url"

	"github.com/feditools/democrablock/internal/http"
)

type FileStore interface {
	http.Module

	GetFile(ctx context.Context, group string, hash []byte, suffix string) ([]byte, error)
	GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, error)
	PutFile(ctx context.Context, group string, hash []byte, suffix string, data []byte) error

	// GetEvidenceFile(ctx context.Context, hash []byte, kind string) ([]byte, error)
	//GetEvidencePresignedURL(ctx context.Context, hash []byte, kind string) (*url.URL, error)
	//PutEvidenceFile(ctx context.Context, hash []byte, kind string, data []byte) error
}

func MakePath(group string, hash []byte) string {
	return fmt.Sprintf(
		"%s/%s",
		group,
		MakeHashDirs(hash),
	)
}

func MakeHashDirs(hash []byte) string {
	return fmt.Sprintf("%x/%x/%x/", hash[0:1], hash[1:2], hash[2:3])
}
