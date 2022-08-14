// Provides configuration for the Storage API
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Application Configuration
type AppConfig struct {
	// Enable DataDog integration
	DataDog       bool   `mapstructure:"data_dog"`
	DataDogApiKey string `mapstructure:"data_dog_api_key"`

	Port        string `mapstructure:"port"`
	AppName     string `mapstructure:"app_name" default:"Storage-API"`
	ServiceName string `mapstructure:"service_name" default:"storage-api"`
	Env         string `mapstructure:"env" default:"dev"`
	Version     string `mapstructure:"version" default:"0.0.1"`
}

func NewAppConfig(path *string, file *string) (AppConfig, error) {
	viper.AddConfigPath(*path)
	viper.SetConfigName(*file)
	viper.SetConfigType("yaml")

	read_err := viper.ReadInConfig()
	if read_err != nil {
		panic(fmt.Errorf("fatal error config file: %w", read_err))
	}

	var config AppConfig
	parse_err := viper.Unmarshal(&config)
	if parse_err != nil {
		panic(fmt.Errorf("cannot parse config %s", parse_err))
	}

	log.Printf("config %+v", config)
	return config, parse_err
}
