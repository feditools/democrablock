package flag

import (
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// Server adds all flags for running the server.
func Server(cmd *cobra.Command, values config.Values) {
	FileStore(cmd, values)
	Redis(cmd, values)

	// server
	cmd.PersistentFlags().String(config.Keys.ServerExternalURL, values.ServerExternalURL, usage.ServerExternalURL)
	cmd.PersistentFlags().String(config.Keys.ServerHTTPBind, values.ServerHTTPBind, usage.ServerHTTPBind)
	cmd.PersistentFlags().Bool(config.Keys.ServerMinifyHTML, values.ServerMinifyHTML, usage.ServerMinifyHTML)
	cmd.PersistentFlags().StringArray(config.Keys.ServerRoles, values.ServerRoles, usage.ServerRoles)

	// webapp
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapCSSURI, values.WebappBootstrapCSSURI, usage.WebappBootstrapCSSURI)
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapCSSIntegrity, values.WebappBootstrapCSSIntegrity, usage.WebappBootstrapCSSIntegrity)
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapJSURI, values.WebappBootstrapJSURI, usage.WebappBootstrapJSURI)
	cmd.PersistentFlags().String(config.Keys.WebappBootstrapJSIntegrity, values.WebappBootstrapJSIntegrity, usage.WebappBootstrapJSIntegrity)
	cmd.PersistentFlags().String(config.Keys.WebappFontAwesomeCSSURI, values.WebappFontAwesomeCSSURI, usage.WebappFontAwesomeCSSURI)
	cmd.PersistentFlags().String(config.Keys.WebappFontAwesomeCSSIntegrity, values.WebappFontAwesomeCSSIntegrity, usage.WebappFontAwesomeCSSIntegrity)
	cmd.PersistentFlags().String(config.Keys.WebappLogoSrcDark, values.WebappLogoSrcDark, usage.WebappLogoSrcDark)
	cmd.PersistentFlags().String(config.Keys.WebappLogoSrcLight, values.WebappLogoSrcLight, usage.WebappLogoSrcLight)
}
