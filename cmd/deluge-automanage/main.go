package main

import (
	"os"

	"github.com/ludviglundgren/deluge-automanage/cmd"
	"github.com/ludviglundgren/deluge-automanage/internal/config"

	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(config.InitConfig)

	var rootCmd = &cobra.Command{
		Use:   "deluge-automanage",
		Short: "Manage Deluge with cli",
		Long: `Manage Deluge from command line.
 
Documentation is available at http://github.com/ludviglundgren/deluge-automanage`,
	}

	// override config
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is $HOME/.config/deluge-automanage/deluge-automanage.toml)")

	// load all commands
	rootCmd.AddCommand(cmd.RunAdd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
