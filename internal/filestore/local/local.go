package local

import (
	nethttp "net/http"
	"os"
	"strings"
	"time"

	"github.com/feditools/democrablock/internal/path"

	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/kv"
	"github.com/spf13/viper"
)

func New(k kv.KV) (*Module, error) {
	l := logger.WithField("func", "New")

	fspath := viper.GetString(config.Keys.FileStorePath)
	if !strings.HasSuffix(fspath, "/") {
		fspath += "/"
	}

	if _, err := os.Stat(fspath); os.IsNotExist(err) {
		l.Debugf("directory '%s' doesn't exist, creating", fspath)
		err = os.Mkdir(fspath, 0755)
		if err != nil {
			l.Errorf("can't create directory '%s': %s", fspath, err.Error())

			return nil, err
		}
	}

	return &Module{
		kv: k,

		path:                   fspath,
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
	return "filestore-local"
}

func (m Module) Route(s *http.Server) error {
	fs := s.PathPrefix(path.Filestore).Subrouter()
	fs.Use(m.middleware)
	fs.NotFoundHandler = m.notFoundHandler()
	fs.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	fs.HandleFunc(path.FilestoreSubFile, m.handleGet).Methods(nethttp.MethodGet)

	return nil
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *http.Server) {
	m.srv = s
}
