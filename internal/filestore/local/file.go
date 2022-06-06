package local

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/feditools/democrablock/internal/path"
	"github.com/google/uuid"

	"github.com/feditools/democrablock/internal/filestore"
)

func (m *Module) GetFile(_ context.Context, group string, hash []byte, suffix string) ([]byte, filestore.Error) {
	l := logger.WithField("func", "GetFile")

	objectFilePath := m.path + filestore.MakeObjectPath(group, hash, suffix)
	data, err := os.ReadFile(objectFilePath)
	if err != nil {
		l.Errorf("reading file: %s", err.Error())

		return nil, m.ProcessError(err)
	}

	return data, nil
}

func (m *Module) GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, filestore.Error) {
	l := logger.WithField("func", "GetPresignedURL")

	objectPath := filestore.MakeObjectPath(group, hash, suffix)

	// check if file exists
	objectFilePath := m.path + objectPath
	if _, err := os.Stat(objectFilePath); os.IsNotExist(err) {
		return nil, filestore.ErrNotFound
	}

	// generate token
	token := uuid.New().String()

	// add token to kv
	if err := m.kv.SetFileStorePresignedURL(ctx, token, objectPath, m.presignedURLExpiration); err != nil {
		l.Errorf("kv set: %s", err.Error())

		return nil, m.ProcessError(err)
	}

	return &url.URL{
		Path:     path.Filestore + "/" + objectPath,
		RawQuery: fmt.Sprintf("token=%s", token),
	}, nil
}

func (m *Module) PutFile(_ context.Context, group string, hash []byte, suffix string, data []byte) filestore.Error {
	l := logger.WithField("func", "PutFile")

	objectPath := m.path + filestore.MakePath(group, hash)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) {
		l.Debugf("directory '%s' doesn't exist, creating", objectPath)
		err = os.MkdirAll(objectPath, 0755)
		if err != nil {
			l.Errorf("can't create directory '%s': %s", objectPath, err.Error())

			return m.ProcessError(err)
		}
	}

	objectFilePath := fmt.Sprintf("%s%x.%s", objectPath, hash, suffix)
	if err := os.WriteFile(objectFilePath, data, 0644); err != nil {
		l.Errorf("can't write file directory '%s': %s", objectFilePath, err.Error())

		return m.ProcessError(err)
	}

	return nil
}
