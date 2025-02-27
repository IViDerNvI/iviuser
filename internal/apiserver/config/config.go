package config

import (
	"github.com/ividernvi/iviuser/internal/pkg/options"
)

type Config struct {
	Options *options.Options
}

var defaultConfig *Config

func newConfig() *Config {
	return &Config{
		Options: options.NewOptions(),
	}
}

func DefaultConfig() *Config {
	if defaultConfig == nil {
		defaultConfig = newConfig()
	}
	return defaultConfig
}
