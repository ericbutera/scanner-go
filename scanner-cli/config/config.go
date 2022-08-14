package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// TODO create service specific entries

type AppConfig struct {
	Region              string
	AccessKey           string `mapstructure:"access_key"`
	SecretKey           string `mapstructure:"secret_key"`
	GCPServiceAccount   string `mapstructure:"gcp_service_account"`
	GCPProjectId        string `mapstructure:"gcp_project_id"`
	AzureSubscriptionId string `mapstructure:"azure_subscription_id"`

	// DataDog profiling
	DataDog       bool   `mapstructure:"data_dog"`
	DataDogApiKey string `mapstructure:"data_dog_api_key"`

	AppName     string `mapstructure:"app_name"`
	ServiceName string `mapstructure:"service_name"`
	Env         string `mapstructure:"env"`
	Version     string `mapstructure:"version"`

	Azure bool `mapstructure:"azure"`
	Aws   bool `mapstructure:"aws"`
	Gcp   bool `mapstructure:"gcp"`

	StorageApiUrl string `mapstructure:"storage_api_url"`
}

func NewAppConfig(path *string) (AppConfig, error) {
	viper.AddConfigPath(*path)
	viper.SetConfigName("app.yaml")
	viper.SetConfigType("yaml")

	viper.SetDefault("storage_api_url", "http://localhost:8080")
	viper.SetDefault("app_name", "Storage-CLI")
	viper.SetDefault("service_name", "storage-cli")
	viper.SetDefault("version", "0.0.1")
	viper.SetDefault("data_dog", false)
	viper.SetDefault("aws", false)
	viper.SetDefault("azure", false)
	viper.SetDefault("gcp", false)

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
