package vfs

import (
	"fmt"

	"github.com/c2fo/vfs/v6"
	"github.com/c2fo/vfs/v6/backend/os"
	"github.com/c2fo/vfs/v6/backend/s3"
	"github.com/c2fo/vfs/v6/vfssimple"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/spf13/viper"
)

func New(d db.DB) (*Module, error) {
	l := logger.WithField("func", "New")

	var files vfs.Location
	var err error

	switch viper.GetString(config.Keys.FileStoreType) {
	case os.Scheme:
		files, err = vfssimple.NewLocation(fmt.Sprintf("file://%s", viper.GetString(config.Keys.FileStorePath)))
		if err != nil {
			l.Errorf("can't open file location: %s", err.Error())

			return nil, err
		}
	case s3.Scheme:
		opts := s3.Options{
			AccessKeyID:     viper.GetString(config.Keys.FileStoreAccessKeyID),
			SecretAccessKey: viper.GetString(config.Keys.FileStoreSecretAccessKey),
		}

		if viper.GetString(config.Keys.FileStoreRegion) != "" {
			opts.Region = viper.GetString(config.Keys.FileStoreRegion)
		}
		if viper.GetString(config.Keys.FileStoreEndpoint) != "" {
			opts.Endpoint = viper.GetString(config.Keys.FileStoreEndpoint)
		}

		fs := s3.NewFileSystem().WithOptions(opts)
		files, err = fs.NewLocation(viper.GetString(config.Keys.FileStoreEndpoint), viper.GetString(config.Keys.FileStoreBucket))
		if err != nil {
			l.Errorf("can't open s3 location: %s", err.Error())

			return nil, err
		}
	default:
		l.Warnf("unknown type: %s", viper.GetString(config.Keys.FileStoreType))

		return nil, fmt.Errorf("unknown type: %s", viper.GetString(config.Keys.FileStoreType))
	}

	return &Module{
		db:    d,
		files: files,
	}, nil
}

type Module struct {
	db    db.DB
	files vfs.Location
}
