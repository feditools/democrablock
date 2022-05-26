package webapp

import (
	"context"
	"encoding/gob"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/fedi"
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
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	htmltemplate "html/template"
	"net/url"
	"strings"
	"sync"
	"time"
)

const SessionMaxAge = 30 * 24 * time.Hour // 30 days

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct {
	db        db.DB
	fedi      *fedi.Module
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	srv       *http.Server
	store     sessions.Store
	templates *htmltemplate.Template
	tokenizer *token.Tokenizer

	logoSrcDark   string
	logoSrcLight  string
	headLinks     []libtemplate.HeadLink
	footerScripts []libtemplate.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

//revive:disable:argument-limit
// New returns a new webapp module.
func New(ctx context.Context, d db.DB, r *redis.Client, fMod *fedi.Module, lMod *language.Module, t *token.Tokenizer, mc metrics.Collector) (*Module, error) {
	l := logger.WithField("func", "New")

	// parse external url.
	externalURL, err := url.Parse(viper.GetString(config.Keys.ServerExternalURL))
	if err != nil {
		l.Errorf("parse external url (%s): %s", viper.GetString(config.Keys.ServerExternalURL), err.Error())

		return nil, err
	}

	// fetch new store.
	store, err := redisstore.NewRedisStore(ctx, r.RedisClient())
	if err != nil {
		l.Errorf("create redis store: %s", err.Error())

		return nil, err
	}

	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   "/",
		Domain: externalURL.Host,
		MaxAge: int(SessionMaxAge.Seconds()),
	})

	// register models for GOB
	gob.Register(http.SessionKey(0))

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
		fedi:      fMod,
		language:  lMod,
		metrics:   mc,
		minify:    m,
		store:     store,
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

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *http.Server) {
	m.srv = s
}
