package local

import (
	"os"
	"strings"
	"time"

	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/kv"
	"github.com/spf13/viper"
)

func New(k kv.KV) (*Module, error) {
	l := logger.WithField("func", "New")

	path := viper.GetString(config.Keys.FileStorePath)
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.Debugf("directory '%s' doesn't exist, creating", path)
		err = os.Mkdir(path, 0755)
		if err != nil {
			l.Errorf("can't create directory '%s': %s", path, err.Error())

			return nil, err
		}
	}

	return &Module{
		kv: k,

		path:                   path,
		presignedURLExpiration: viper.GetDuration(config.Keys.FileStorePresignedURLExpiration),
	}, nil
}

type Module struct {
	kv  kv.KV
	srv *http.Server

	path                   string
	presignedURLExpiration time.Duration
}

func (*Module) Name() string {
	return "filestore-minio"
}

func (m Module) Route(s *http.Server) error {
	return nil
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *http.Server) {
	m.srv = s
}
