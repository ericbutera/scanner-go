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

	viper.SetDefault("jaeger", false)
	viper.SetDefault("data_dog", false)
	viper.SetDefault("data_dog_api_key", "")
	viper.SetDefault("app_name", "Storage-API")
	viper.SetDefault("service_name", "storage-api")
	viper.SetDefault("env", "dev")
	viper.SetDefault("version", "0.0.1")

	read_err := viper.ReadInConfig()
	if read_err != nil {
		log.Print(fmt.Errorf("fatal error config file: %w", read_err))
	}

	var config AppConfig
	parse_err := viper.Unmarshal(&config)
	if parse_err != nil {
		log.Print(fmt.Errorf("cannot parse config %s", parse_err))
	}

	log.Printf("config %+v", config)
	return config, parse_err
}
