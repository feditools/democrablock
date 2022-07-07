package minio

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/feditools/democrablock/internal/filestore"
	libhttp "github.com/feditools/go-lib/http"
	"github.com/minio/minio-go/v7"
)

func (m *Module) GetFile(ctx context.Context, group string, hash []byte, suffix string) ([]byte, filestore.Error) {
	l := logger.WithField("func", "GetFile")

	objectPath := filestore.MakeObjectPath(group, hash, suffix)

	reader, err := m.mc.GetObject(ctx, m.bucket, objectPath, minio.GetObjectOptions{})
	if err != nil {
		l.Errorf("can't get image data: %s", err.Error())

		return nil, m.ProcessError(err)
	}
	defer reader.Close()

	var b []byte
	_, err = reader.Read(b)
	if err != nil {
		l.Errorf("can't read image data: %s", err.Error())

		return nil, m.ProcessError(err)
	}

	return b, nil
}

func (m *Module) GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, filestore.Error) {
	l := logger.WithField("func", "GetPresignedURL")

	objectPath := filestore.MakeObjectPath(group, hash, suffix)

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%x.%s\"", hash, suffix))

	presignedURL, err := m.mc.PresignedGetObject(ctx, m.bucket, objectPath, m.presignedURLExpiration, reqParams)
	if err != nil {
		l.Errorf("getting prosigned url %s:%s: %s", m.bucket, objectPath, err.Error())

		return nil, err
	}

	return presignedURL, nil
}

func (m *Module) PutFile(ctx context.Context, group string, hash []byte, suffix string, data []byte) filestore.Error {
	l := logger.WithField("func", "PutFile")

	objectPath := filestore.MakeObjectPath(group, hash, suffix)
	contentType := string(libhttp.ToMime(libhttp.Suffix(suffix)))

	reader := bytes.NewReader(data)
	n, err := m.mc.PutObject(ctx, m.bucket, objectPath, reader, reader.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		l.Errorf("can't put image to %s:%s: %s", m.bucket, objectPath, err.Error())

		return err
	}
	l.Debugf("wrote %d bytes to %s:%s", n.Size, m.bucket, objectPath)

	return nil
}
