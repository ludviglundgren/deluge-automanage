package cmd

import (
	"fmt"
	"os"

	"github.com/ludviglundgren/deluge-automanage/internal/domain"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	Config  domain.AppConfig
	rootCmd = &cobra.Command{
		Version: "v0.1",
		Use:     "deluge-automanage",
		Short:   "Manage Deluge with cli",
		Long: `Managa Deluge from command line.
 
Documentation is available at http://github.com/ludviglundgren/deluge-automanage`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/deluge-automanage/deluge-automanage.toml)")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in directories
		viper.SetConfigName(".deluge-automanage")
		viper.AddConfigPath(home)
		viper.AddConfigPath("$HOME/.config/deluge-automanage") // call multiple times to add many search paths
		viper.AddConfigPath(".")                               // optionally look for config in the working directory
	}

	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Could not read config file:", viper.ConfigFileUsed())
		os.Exit(1)
	}
	// fmt.Println("Using config file:", viper.ConfigFileUsed())

	err := viper.Unmarshal(&Config)
	if err != nil {
		// fmt.Fatalf("unable to decode into struct, %v", err)
		os.Exit(1)
	}

	// fmt.Printf("config: %v\n", &Config)
}
