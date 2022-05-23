package webapp

import (
	"context"
	"encoding/gob"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/db"
	"github.com/feditools/democrablock/internal/grpc"
	"github.com/feditools/democrablock/internal/http"
	"github.com/feditools/democrablock/internal/http/template"
	"github.com/feditools/democrablock/internal/kv"
	"github.com/feditools/democrablock/internal/kv/redis"
	"github.com/feditools/democrablock/internal/metrics"
	"github.com/feditools/democrablock/internal/path"
	"github.com/feditools/democrablock/internal/token"
	"github.com/feditools/go-lib/language"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/login/pkg/oauth"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"golang.org/x/oauth2"
	htmltemplate "html/template"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct {
	db        db.DB
	grpc      *grpc.Client
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	oauth     *oauth.Client
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

const ThirtyDays = 30 * 24 * time.Hour

//revive:disable:argument-limit
// New returns a new webapp module.
func New(ctx context.Context, d db.DB, g *grpc.Client, r *redis.Client, lMod *language.Module, t *token.Tokenizer, mc metrics.Collector) (http.Module, error) {
	l := logger.WithField("func", "New")

	oauthConfig := &oauth.Config{
		CallbackURL:  viper.GetString(config.Keys.ServerExternalURL) + path.CallbackOauth,
		ServerURL:    viper.GetString(config.Keys.OAuthServerURL),
		ClientID:     viper.GetString(config.Keys.OAuthClientID),
		ClientSecret: viper.GetString(config.Keys.OAuthClientSecret),
	}
	oauthClient, err := oauth.New(ctx, oauthConfig)
	if err != nil {
		l.Errorf("create oauth client: %s", err.Error())

		return nil, err
	}

	// Fetch new store.
	store, err := redisstore.NewRedisStore(ctx, r.RedisClient())
	if err != nil {
		l.Errorf("create redis store: %s", err.Error())

		return nil, err
	}

	serverExternalURL, err := url.Parse(viper.GetString(config.Keys.ServerExternalURL))
	if err != nil {
		l.Errorf("parsing external url: %s", err.Error())

		return nil, err
	}
	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   "/",
		Domain: serverExternalURL.Host,
		MaxAge: int(ThirtyDays.Seconds()),
	})

	// Register models for GOB
	gob.Register(SessionKey(0))
	gob.Register(oauth2.Token{})

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
		grpc:      g,
		language:  lMod,
		metrics:   mc,
		minify:    m,
		oauth:     oauthClient,
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
