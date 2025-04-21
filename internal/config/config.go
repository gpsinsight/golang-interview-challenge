package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment string `envconfig:"ENVIRONMENT"`
	Postgres    PostgresConfig
	Kafka       KafkaConfig[TopicConfig]
}

// New creates a new Config from environment variables and env files in a specified directory
func New() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
