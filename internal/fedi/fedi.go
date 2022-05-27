package fedi

import (
	"context"
	"net/url"

	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/go-lib/fedihelper"
	"github.com/feditools/go-lib/fedihelper/mastodon"
	"github.com/spf13/viper"
)

func New(d db.DB, h fedihelper.HTTP, k fedihelper.KV, t *token.Tokenizer) (*Module, error) {
	appName := viper.GetString(config.Keys.ApplicationName)
	appWebsite := viper.GetString(config.Keys.ApplicationWebsite)
	externalURL := viper.GetString(config.Keys.ServerExternalURL)

	// prep fedi helpers
	var fediHelpers []fedihelper.Helper
	mastoHelper, err := mastodon.New(k, appName, appWebsite, externalURL)
	if err != nil {
		return nil, err
	}
	fediHelpers = append(fediHelpers, mastoHelper)

	// prep fedi
	newModule := &Module{
		db:   d,
		tokz: t,
	}

	fedi, err := fedihelper.New(h, k, appName, fediHelpers)
	if err != nil {
		return nil, err
	}

	fedi.SetCreateAccountHandler(newModule.CreateAccountHandler)
	fedi.SetCreateInstanceHandler(newModule.CreateInstanceHandler)
	fedi.SetGetAccountHandler(newModule.GetAccountHandler)
	fedi.SetGetInstanceHandler(newModule.GetInstanceHandler)
	fedi.SetGetTokenHandler(newModule.GetTokenHandler)
	fedi.SetNewAccountHandler(newModule.NewAccountHandler)
	fedi.SetNewInstanceHandler(newModule.NewInstanceHandler)
	fedi.SetUpdateInstanceHandler(newModule.UpdateInstanceHandler)

	newModule.helper = fedi

	return newModule, nil
}

type Module struct {
	db   db.DB
	tokz *token.Tokenizer

	helper *fedihelper.FediHelper
}

func (m *Module) GetLoginURL(ctx context.Context, act string) (*url.URL, error) {
	return m.helper.GetLoginURL(ctx, act)
}

func (m *Module) Helper(s fedihelper.Software) fedihelper.Helper {
	return m.helper.Helper(s)
}
