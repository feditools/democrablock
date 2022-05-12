package flag

import (
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// GRPC adds all flags for running the grpc client.
func GRPC(cmd *cobra.Command, values config.Values) {
	cmd.PersistentFlags().String(config.Keys.GRPCLoginAddress, values.GRPCLoginAddress, usage.GRPCLoginAddress)
	cmd.PersistentFlags().String(config.Keys.GRPCLoginToken, values.GRPCLoginToken, usage.GRPCLoginToken)
}
