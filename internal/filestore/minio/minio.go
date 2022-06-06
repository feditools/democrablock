package minio

import (
	"time"

	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/http"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func New() (*Module, error) {
	l := logger.WithField("func", "New")

	endpoint := viper.GetString(config.Keys.FileStoreEndpoint)
	accessKeyID := viper.GetString(config.Keys.FileStoreAccessKeyID)
	secretAccessKey := viper.GetString(config.Keys.FileStoreSecretAccessKey)
	useSSL := viper.GetBool(config.Keys.FileStoreUseTLS)

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		l.Errorf("creating minio client: %s", err.Error())

		return nil, err
	}

	return &Module{
		mc:                     mc,
		bucket:                 viper.GetString(config.Keys.FileStoreBucket),
		presignedURLExpiration: viper.GetDuration(config.Keys.FileStorePresignedURLExpiration),
	}, nil
}

type Module struct {
	mc                     *minio.Client
	bucket                 string
	presignedURLExpiration time.Duration
}

func (m *Module) Name() string {
	return "filestore-minio"
}

func (m *Module) Route(_ *http.Server) error {
	return nil
}

func (m *Module) SetServer(_ *http.Server) {}
