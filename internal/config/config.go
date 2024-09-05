package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Opensearch OpenSearchConfig
}

type OpenSearchConfig struct {
	Clusters []string `mapstructure:"clusters"`
	RoleARN  string   `mapstructure:"role_arn"`
	Region   string   `mapstructure:"region"`
	Bucket   string   `mapstructure:"bucket"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName(os.Getenv("ENV"))
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

func NewOpensearchConfig(logger *zap.Logger) (*Config, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	logger.Info("Config loaded successfully")
	return cfg, nil
}
