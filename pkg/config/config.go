package config

import (
	"os"
	"time"
)

type Config struct {
	AppPort string
	MainDB  Database
}

func NewConfig() *Config {
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		MainDB: Database{
			Host:              os.Getenv("DB_POSTGRES_HOST"),
			Port:              EnvAsInt(os.Getenv("DB_POSTGRES_PORT"), 5432),
			User:              os.Getenv("DB_POSTGRES_USERNAME"),
			Password:          os.Getenv("DB_POSTGRES_PASSWORD"),
			DatabaseName:      os.Getenv("DB_POSTGRES_DATABASE"),
			MaxIdleConnection: EnvAsInt("DB_POSTGRES_MAX_IDLE_CONNECTION", 25),
			MaxIdleTime:       EnvAsDuration("DB_POSTGRES_MAX_IDLE_TIME", 5*time.Minute),
			MaxOpenConnection: EnvAsInt("DB_POSTGRES_MAX_OPEN_CONNECTION", 25),
			MaxLifetime:       EnvAsDuration("DB_POSTGRES_MAX_LIFETIME", 5*time.Minute),
		},
	}
}
