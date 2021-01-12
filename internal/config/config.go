package config

import (
	"fmt"
	"os"

	"github.com/ludviglundgren/deluge-automanage/internal/domain"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	CfgFile string
	Config  domain.AppConfig
	Deluge  domain.DelugeConfig
	Rules   domain.Rules
)

func InitConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			//er(err)
			fmt.Println("Could not read home dir:", err)
			os.Exit(1)
		}

		// Search config in directories
		viper.SetConfigName(".deluge-automanage")
		viper.AddConfigPath(home)
		viper.AddConfigPath("$HOME/.config/deluge-automanage") // call multiple times to add many search paths
		viper.AddConfigPath(".")                               // optionally look for config in the working directory
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Could not read config file:", viper.ConfigFileUsed())
		os.Exit(1)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		// fmt.Fatalf("unable to decode into struct, %v", err)
		os.Exit(1)
	}

	Deluge = Config.Deluge
	Rules = Config.Rules
}
