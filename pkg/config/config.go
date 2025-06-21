package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
}

func NewConfig() *Config {
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
	}
}

func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", filename, err)
	}
}
