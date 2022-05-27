package main

import (
	"context"
	"fmt"

	"github.com/feditools/democrablock/cmd/democrablock/action"
	"github.com/feditools/democrablock/cmd/democrablock/flag"
	"github.com/feditools/democrablock/internal/config"
	"github.com/feditools/democrablock/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Version is the software version.
var Version string

// Commit is the git commit.
var Commit string

func main() {
	l := logger.WithField("func", "main")

	var v string
	if len(Commit) < 7 {
		v = Version
	} else {
		v = Version + " " + Commit[:7]
	}

	// set software version
	viper.Set(config.Keys.SoftwareVersion, v)

	rootCmd := &cobra.Command{
		Use:   "democrablock",
		Short: "democrablock - fediverse democrablock server",

		Version:       v,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	flag.Global(rootCmd, config.Defaults)

	err := viper.BindPFlag(config.Keys.ConfigPath, rootCmd.PersistentFlags().Lookup(config.Keys.ConfigPath))
	if err != nil {
		l.Fatalf("error binding config flag: %s", err)

		return
	}

	// add commands
	rootCmd.AddCommand(accountCommands())
	rootCmd.AddCommand(databaseCommands())
	rootCmd.AddCommand(serverCommands())

	err = rootCmd.Execute()
	if err != nil {
		l.Fatalf("error executing command: %s", err)
	}
}

func preRun(cmd *cobra.Command) error {
	if err := config.Init(cmd.Flags()); err != nil {
		return fmt.Errorf("error initializing config: %s", err)
	}

	if err := config.ReadConfigFile(); err != nil {
		return fmt.Errorf("error reading config: %s", err)
	}

	return nil
}

func run(ctx context.Context, act action.Action) error {
	if err := log.Init(); err != nil {
		return fmt.Errorf("error initializing log: %s", err)
	}

	return act(ctx)
}
