package flag

import (
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// CouncilInit adds all flags for running the council init command.
func CouncilInit(cmd *cobra.Command, values config.Values) {
	Redis(cmd, values)

	cmd.PersistentFlags().StringSlice(config.Keys.CouncilMembers, values.CouncilMembers, usage.CouncilMembers)
}
