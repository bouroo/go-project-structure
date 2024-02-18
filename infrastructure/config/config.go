package config

import (
	"sync"

	"github.com/spf13/viper"
)

type viperConfig struct {
	mutex      sync.RWMutex
	configName string
	viper      *viper.Viper
}

type AppConfig interface {
	GetViper() *viper.Viper
	LoadConfig() (err error)
	CheckConfigs(keysUsed map[string]bool) (err error)
	WatchConfig() (err error)
}

// NewViperConfig creates a new AppConfig using Viper configuration.
//
// It takes a configName string as a parameter and returns an AppConfig.
func NewViperConfig(configName string) AppConfig {
	return &viperConfig{
		mutex:      sync.RWMutex{},
		configName: configName,
		viper:      viper.New(),
	}
}
