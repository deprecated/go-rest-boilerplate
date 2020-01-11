package config

import (
	"fmt"
	
	"github.com/spf13/viper"

	"github.com/deprecated/go-rest-boilerplate/models"
)

// Config is global object that holds all application level variables.
var Config models.Configurations

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.SetEnvPrefix("boilerplate")
	v.AutomaticEnv()

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	return v.Unmarshal(&Config)
}