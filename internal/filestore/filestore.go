package filestore

import (
	"context"
	"fmt"
	"net/url"

	"github.com/feditools/democrablock/internal/http"
)

type FileStore interface {
	http.Module

	GetFile(ctx context.Context, group string, hash []byte, suffix string) ([]byte, Error)
	GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, Error)
	PutFile(ctx context.Context, group string, hash []byte, suffix string, data []byte) Error
}

func MakeObjectPath(group string, hash []byte, suffix string) string {
	return fmt.Sprintf(
		"%s%x.%s",
		MakePath(group, hash),
		hash,
		suffix,
	)
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
