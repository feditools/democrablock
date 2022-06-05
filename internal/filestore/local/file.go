package local

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/feditools/democrablock/internal/path"
	"github.com/google/uuid"

	"github.com/feditools/democrablock/internal/filestore"
)

func (m *Module) GetFile(_ context.Context, group string, hash []byte, suffix string) ([]byte, error) {
	l := logger.WithField("func", "GetFile")

	objectFilePath := fmt.Sprintf("%s%s%x.%s", m.path, filestore.MakePath(group, hash), hash, suffix)
	data, err := os.ReadFile(objectFilePath)
	if err != nil {
		l.Errorf("reading file: %s", err.Error())

		return nil, err
	}

	return data, nil
}

func (m *Module) GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, error) {
	l := logger.WithField("func", "GetPresignedURL")

	objectPath := fmt.Sprintf("%s%x.%s", filestore.MakePath(group, hash), hash, suffix)

	// check if file exists
	objectFilePath := m.path + objectPath
	if _, err := os.Stat(objectFilePath); os.IsNotExist(err) {
		return nil, errors.New("file not found")
	}

	// generate token
	token := uuid.New().String()

	// add token to kv
	if err := m.kv.SetFileStorePresignedURL(ctx, token, objectPath, m.presignedURLExpiration); err != nil {
		l.Errorf("kv set: %s", err.Error())

		return nil, err
	}

	return &url.URL{
		Path:     path.Filestore + "/" + objectPath,
		RawQuery: fmt.Sprintf("token=%s", token),
	}, nil
}

func (m *Module) PutFile(_ context.Context, group string, hash []byte, suffix string, data []byte) error {
	l := logger.WithField("func", "PutFile")

	objectPath := m.path + filestore.MakePath(group, hash)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) {
		l.Debugf("directory '%s' doesn't exist, creating", objectPath)
		err = os.MkdirAll(objectPath, 0755)
		if err != nil {
			l.Errorf("can't create directory '%s': %s", objectPath, err.Error())

			return err
		}
	}

	objectFilePath := fmt.Sprintf("%s%x.%s", objectPath, hash, suffix)
	if err := os.WriteFile(objectFilePath, data, 0644); err != nil {
		l.Errorf("can't write file directory '%s': %s", objectFilePath, err.Error())

		return err
	}

	return nil
}
