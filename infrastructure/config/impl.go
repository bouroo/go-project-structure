package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// GetViper returns the Viper instance from the viperConfig.
//
// No parameters.
// Returns a pointer to viper.Viper.
func (v *viperConfig) GetViper() *viper.Viper {
	return v.viper
}

// LoadConfig loads the configuration for the viperConfig.
//
// It sets default values for the configuration and then loads the common and environment-specific
// configuration files. It also overrides the config file with environment variables. Returns an error if
// any operation fails.
func (v *viperConfig) LoadConfig(configName string) (err error) {
	if v.viper == nil {
		v.viper = viper.New()
	}

	if len(configName) != 0 {
		v.configName = configName
	} else {
		v.configName = "local"
	}

	// set default values
	v.viper.SetDefault("app.listen", "localhost")
	v.viper.SetDefault("app.port.http", "8080")
	v.viper.SetDefault("app.debug", false)

	if _, err := os.Stat(v.configPath); err == nil {
		// add config files path
		v.viper.AddConfigPath(v.configPath)

		// REQUIRED if the config file does not have the extension in the name
		v.viper.SetConfigType("toml")

		// load common config
		v.viper.SetConfigName("common")

		// find and read common config file
		v.viper.ReadInConfig()

		if _, err := os.Stat(v.configName); err == nil {
			// name of config file (without extension)
			v.viper.SetConfigName(v.configName)

			// Find and read the config file
			if err = v.viper.MergeInConfig(); err != nil {
				return err
			}
		}
	}

	// override config file with environment variables
	v.viper.AutomaticEnv()
	// convert ENV from APP_NAME to viper key app.name
	v.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return
}

// CheckConfigs checks the given keys against the viper configuration and returns an error if any key is missing.
//
// It takes a map of keysUsed which contains the keys to be checked. It returns an error if any key is missing.
// Returns an error.
func (v *viperConfig) CheckConfigs(keysUsed map[string]bool) (err error) {
	if len(keysUsed) == 0 {
		return
	}
	var missingKeys []string
	for key, required := range keysUsed {
		if !required {
			continue
		} else if v.viper.Get(key) == nil {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		err = errors.New("missing config keys: " + strings.Join(missingKeys, ", "))
		return
	}

	return
}

// WatchConfig watches the configuration and reloads it on change.
//
// No parameters.
// Returns an error.
func (v *viperConfig) WatchConfig() (err error) {
	v.viper.WatchConfig()
	v.viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		v.mutex.Lock()
		defer v.mutex.Unlock()
		err := v.LoadConfig(v.configName)
		if err != nil {
			return
		}
	})
	return
}
