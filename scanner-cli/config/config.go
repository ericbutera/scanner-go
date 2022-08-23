package config

import (
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

func Load(path string, conf *AppConfig) error {
	Defaults()

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(conf); err != nil {
		return err
	}

	return nil
}

func NewAppConfig(path string) (*AppConfig, error) {
	Defaults()

	config := &AppConfig{}
	if err := Load(path, config); err != nil {
		return nil, err
	}

	return config, nil
}

func Defaults() {
	viper.SetDefault("storage_api_url", "http://localhost:8080")
	viper.SetDefault("app_name", "Storage-CLI")
	viper.SetDefault("service_name", "storage-cli")
	viper.SetDefault("version", "0.0.1")
	viper.SetDefault("data_dog", false)
	viper.SetDefault("aws", false)
	viper.SetDefault("azure", false)
	viper.SetDefault("gcp", false)
}
