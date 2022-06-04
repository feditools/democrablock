package local

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/feditools/democrablock/internal/filestore"
)

func (m *Module) GetFile(ctx context.Context, group string, hash []byte, suffix string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (m *Module) GetPresignedURL(ctx context.Context, group string, hash []byte, suffix string) (*url.URL, error) {
	// TODO implement me
	panic("implement me")
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
