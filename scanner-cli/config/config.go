package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Region              string
	AccessKey           string `mapstructure:"access_key"`
	SecretKey           string `mapstructure:"secret_key"`
	GCPServiceAccount   string `mapstructure:"gcp_service_account"`
	AzureSubscriptionId string `mapstructure:"azure_subscription_id"`

	DataDog     bool   `mapstructure:"data_dog" default:"false"`
	AppName     string `mapstructure:"app_name" default:"Storage-CLI"`
	ServiceName string `mapstructure:"service_name" default:"storage-cli"`
	Env         string `mapstructure:"env" default:"dev"`
	Version     string `mapstructure:"version" default:"0.0.1"`
}

func NewAppConfig(path *string) (AppConfig, error) {
	viper.AddConfigPath(*path)
	viper.SetConfigName("app.yaml")
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
	return config, parse_err
}
