package config

import (
	"sync"

	"github.com/spf13/viper"
)

type viperConfig struct {
	mutex      sync.RWMutex
	configPath string
	configName string
	viper      *viper.Viper
}

type AppConfig interface {
	GetViper() *viper.Viper
	LoadConfig(configName string) (err error)
	CheckConfigs(keysUsed map[string]bool) (err error)
	WatchConfig() (err error)
}

// NewAppConfig creates a new AppConfig using Viper configuration.
//
// It takes a configName string as a parameter and returns an AppConfig.
func NewAppConfig(configPath string) AppConfig {
	return &viperConfig{
		mutex:      sync.RWMutex{},
		configPath: configPath,
		viper:      viper.New(),
	}
}
