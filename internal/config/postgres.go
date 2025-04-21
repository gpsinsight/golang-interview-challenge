package config

import "fmt"

type PostgresConfig struct {
	Password     string `envconfig:"POSTGRES_PASSWORD"`
	User         string `envconfig:"POSTGRES_USER"     default:"postgres"`
	Port         int    `envconfig:"POSTGRES_PORT"     default:"5432"`
	Database     string `envconfig:"POSTGRES_DB"       default:"postgres"`
	Host         string `envconfig:"POSTGRES_HOST"`
	SSLMode      string `envconfig:"POSTGRES_SSL_MODE" default:"require"`
	MaxOpenConns int    `envconfig:"POSTGRES_MAX_OPEN_CONNS"`
}

// ConnectionString builds a postgres connection string from the configured values
func (pc PostgresConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pc.Host,
		pc.Port,
		pc.User,
		pc.Password,
		pc.Database,
		pc.SSLMode,
	)
}
