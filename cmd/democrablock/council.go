package main

import (
	"github.com/feditools/democrablock/cmd/democrablock/action/council"
	"github.com/feditools/democrablock/cmd/democrablock/flag"
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// councilCommands returns the 'council' subcommand.
func councilCommands() *cobra.Command {
	councilCmd := &cobra.Command{
		Use:   "council",
		Short: "manage the council",
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "initialize the council",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), council.Init)
		},
	}
	flag.CouncilInit(initCmd, config.Defaults)
	councilCmd.AddCommand(initCmd)

	return councilCmd
}
