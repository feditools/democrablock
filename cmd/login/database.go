package main

import (
	"github.com/feditools/democrablock/cmd/login/action/database"
	"github.com/feditools/democrablock/cmd/login/flag"
	"github.com/feditools/democrablock/internal/config"
	"github.com/spf13/cobra"
)

// databaseCommands returns the 'database' subcommand.
func databaseCommands() *cobra.Command {
	databaseCmd := &cobra.Command{
		Use:   "database",
		Short: "manage the database",
	}
	flag.Database(databaseCmd, config.Defaults)

	databaseMigrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "run migrations",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), database.Migrate)
		},
	}
	flag.DatabaseMigrate(databaseMigrateCmd, config.Defaults)
	databaseCmd.AddCommand(databaseMigrateCmd)

	return databaseCmd
}
