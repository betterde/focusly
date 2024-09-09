package config

import (
	"github.com/betterde/focusly/internal/journal"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Conf *Config

type Config struct {
	// Add config item to here
	Env      string   `yaml:"env"`
	HTTP     HTTP     `yaml:"http"`
	Logging  Logging  `yaml:"logging"`
	Database Database `yaml:"database"`
}

type HTTP struct {
	Listen  string `yaml:"listen"`
	TLSKey  string `yaml:"tlsKey"`
	TLSCert string `yaml:"tlsCert"`
}

type Logging struct {
	Level string `yaml:"level"`
}

type Database struct {
	URL       string `yaml:"url"`
	Database  string `yaml:"database"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Namespace string `yaml:"namespace"`
}

func Parse(file string, envPrefix string) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".focusly")
	}

	// read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(envPrefix)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		journal.Logger.Errorf("Failed to read configuration file: %s", err)
		os.Exit(1)
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	err := viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
