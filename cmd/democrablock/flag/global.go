package flag

import (
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// Global adds flags that are common to all commands.
func Global(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.ConfigPath, values.ConfigPath, usage.ConfigPath)
	cmd.PersistentFlags().String(config.Keys.LogLevel, values.LogLevel, usage.LogLevel)

	// application
	cmd.PersistentFlags().String(config.Keys.ApplicationName, values.ApplicationName, usage.ApplicationName)
	cmd.PersistentFlags().String(config.Keys.EncryptionKey, values.EncryptionKey, usage.EncryptionKey)

	// database
	cmd.PersistentFlags().String(config.Keys.DBType, values.DBType, usage.DBType)
	cmd.PersistentFlags().String(config.Keys.DBAddress, values.DBAddress, usage.DBAddress)
	cmd.PersistentFlags().Int(config.Keys.DBPort, values.DBPort, usage.DBPort)
	cmd.PersistentFlags().String(config.Keys.DBUser, values.DBUser, usage.DBUser)
	cmd.PersistentFlags().String(config.Keys.DBPassword, values.DBPassword, usage.DBPassword)
	cmd.PersistentFlags().String(config.Keys.DBDatabase, values.DBDatabase, usage.DBDatabase)
	cmd.PersistentFlags().String(config.Keys.DBTLSMode, values.DBTLSMode, usage.DBTLSMode)
	cmd.PersistentFlags().String(config.Keys.DBTLSCACert, values.DBTLSCACert, usage.DBTLSCACert)

	// metrics
	cmd.PersistentFlags().String(config.Keys.MetricsStatsDAddress, values.MetricsStatsDAddress, usage.MetricsStatsDAddress)
	cmd.PersistentFlags().String(config.Keys.MetricsStatsDPrefix, values.MetricsStatsDPrefix, usage.MetricsStatsDPrefix)
}
