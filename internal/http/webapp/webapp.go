package webapp

import (
	"context"
	"encoding/gob"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/democrablock/internal/kv"
	"github.com/feditools/democrablock/internal/kv/redis"
	"github.com/feditools/democrablock/internal/metrics"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	minify "github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	htmltemplate "html/template"
	"strings"
	"sync"
	"time"
)

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct {
	db        db.DB
	store     sessions.Store
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	templates *htmltemplate.Template
	tokenizer *token.Tokenizer

	logoSrcDark   string
	logoSrcLight  string
	headLinks     []libtemplate.HeadLink
	footerScripts []libtemplate.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

const ThirtyDays = 30 * 24 * time.Hour

//revive:disable:argument-limit
// New returns a new webapp module.
func New(ctx context.Context, d db.DB, r *redis.Client, lMod *language.Module, t *token.Tokenizer, mc metrics.Collector) (http.Module, error) {
	l := logger.WithField("func", "New")

	// Fetch new store.
	store, err := redisstore.NewRedisStore(ctx, r.RedisClient())
	if err != nil {
		l.Errorf("create redis store: %s", err.Error())

		return nil, err
	}

	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   "/",
		Domain: viper.GetString(config.Keys.ServerExternalHostname),
		MaxAge: int(ThirtyDays.Seconds()),
	})

	// Register models for GOB
	gob.Register(SessionKey(0))

	// minify
	var m *minify.M
	if viper.GetBool(config.Keys.ServerMinifyHTML) {
		m = minify.New()
		m.AddFunc("text/html", html.Minify)
	}

	// get templates
	tmpl, err := template.New(t)
	if err != nil {
		l.Errorf("create temates: %s", err.Error())

		return nil, err
	}

	// generate head links
	hl := []libtemplate.HeadLink{
		{
			HRef:        viper.GetString(config.Keys.WebappBootstrapCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: COAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappBootstrapCSSIntegrity),
		},
		{
			HRef:        viper.GetString(config.Keys.WebappFontAwesomeCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: COAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappFontAwesomeCSSIntegrity),
		},
	}
	paths := []string{
		path.FileDefaultCSS,
	}
	for _, p := range paths {
		signature, err := getSignature(strings.TrimPrefix(p, "/"))
		if err != nil {
			l.Errorf("getting signature for %s: %s", p, err.Error())
		}

		hl = append(hl, libtemplate.HeadLink{
			HRef:        p,
			Rel:         "stylesheet",
			CrossOrigin: COAnonymous,
			Integrity:   signature,
		})
	}

	// generate head links
	fs := []libtemplate.Script{
		{
			Src:         viper.GetString(config.Keys.WebappBootstrapJSURI),
			CrossOrigin: COAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappBootstrapJSIntegrity),
		},
	}

	return &Module{
		db:        d,
		store:     store,
		language:  lMod,
		metrics:   mc,
		minify:    m,
		templates: tmpl,
		tokenizer: t,

		logoSrcDark:   viper.GetString(config.Keys.WebappLogoSrcDark),
		logoSrcLight:  viper.GetString(config.Keys.WebappLogoSrcLight),
		headLinks:     hl,
		footerScripts: fs,

		sigCache: map[string]string{},
	}, nil
} //revive:enable:argument-limit

// Name return the module name.
func (*Module) Name() string {
	return config.ServerRoleWebapp
}