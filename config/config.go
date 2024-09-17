package config

import (
	"errors"
	"fmt"
	"github.com/betterde/focusly/internal/journal"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	keys = []string{
		"ENV",
		"HTTP.LISTEN",
		"HTTP.TLSKEY",
		"HTTP.TLSCERT",
		"GRPC.LISTEN",
		"GRPC.TLSKEY",
		"GRPC.TLSCERT",
		"LOGGING.LEVEL",
		"DATABASE.URL",
		"DATABASE.DATABASE",
		"DATABASE.USERNAME",
		"DATABASE.PASSWORD",
		"DATABASE.NAMESPACE",
	}
	Conf *Config
)

type Config struct {
	// Add config item to here
	Env      string   `yaml:"env" mapstructure:"ENV"`
	HTTP     HTTP     `yaml:"http" mapstructure:"HTTP"`
	GRPC     GRPC     `yaml:"grpc" mapstructure:"GRPC"`
	Logging  Logging  `yaml:"logging" mapstructure:"LOGGING"`
	Database Database `yaml:"database" mapstructure:"DATABASE"`
}

type HTTP struct {
	Listen  string `yaml:"listen" mapstructure:"LISTEN"`
	TLSKey  string `yaml:"tlsKey" mapstructure:"TLSKEY"`
	TLSCert string `yaml:"tlsCert" mapstructure:"TLSCERT"`
}

type GRPC struct {
	Listen  string `yaml:"listen" mapstructure:"LISTEN"`
	TLSKey  string `yaml:"tlsKey" mapstructure:"TLSKEY"`
	TLSCert string `yaml:"tlsCert" mapstructure:"TLSCERT"`
}

type Logging struct {
	Level string `yaml:"level" mapstructure:"LEVEL"`
}

type Database struct {
	URL       string `yaml:"url" mapstructure:"URL"`
	Database  string `yaml:"database" mapstructure:"DATABASE"`
	Username  string `yaml:"username" mapstructure:"USERNAME"`
	Password  string `yaml:"password" mapstructure:"PASSWORD"`
	Namespace string `yaml:"namespace" mapstructure:"NAMESPACE"`
}

func Parse(file string, envPrefix string) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".focusly")
		viper.AddConfigPath("/etc/focusly")
	}

	replacer := strings.NewReplacer(".", "_")

	// read in environment variables that match
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(envPrefix)

	var notFoundError viper.ConfigFileNotFoundError

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil && errors.As(err, &notFoundError) {
		for _, key := range keys {
			if err := viper.BindEnv(key, fmt.Sprintf("%s_%s", envPrefix, replacer.Replace(key))); err != nil {
				journal.Logger.Error(err)
			}
		}
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	err := viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
