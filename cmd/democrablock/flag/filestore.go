package flag

import (
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// FileStore adds flags that are common to filestore.
func FileStore(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.FileStoreType, values.FileStoreType, usage.FileStoreType)
	cmd.PersistentFlags().String(config.Keys.FileStorePath, values.FileStorePath, usage.FileStorePath)
	cmd.PersistentFlags().String(config.Keys.FileStoreBucket, values.FileStoreBucket, usage.FileStoreBucket)
	cmd.PersistentFlags().String(config.Keys.FileStoreEndpoint, values.FileStoreEndpoint, usage.FileStoreEndpoint)
	cmd.PersistentFlags().String(config.Keys.FileStoreRegion, values.FileStoreRegion, usage.FileStoreRegion)
	cmd.PersistentFlags().Bool(config.Keys.FileStoreUseTLS, values.FileStoreUseTLS, usage.FileStoreUseTLS)
	cmd.PersistentFlags().Duration(config.Keys.FileStorePresignedURLExpiration, values.FileStorePresignedURLExpiration, usage.FileStorePresignedURLExpiration)
}
