package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Opensearch OpenSearchConfig
}

type OpenSearchConfig struct {
	Clusters []string `mapstructure:"clusters"`
	RoleARN  string   `mapstructure:"role_arn"`
	Region   string
	Bucket   string
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("dev")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}

func NewOpensearchConfig() (*Config, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
