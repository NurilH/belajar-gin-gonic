package config

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

const defaultConfigFile = "config.env"

func LoadEnv(configFile string) {
	if configFile == "" {
		configFile = defaultConfigFile
	}
	if err := godotenv.Load(configFile); err != nil {
		log.Fatalf("Error loading %s file \n", configFile)
	}
}

func Env(key string) string {
	envMap, errReadEnv := godotenv.Read(defaultConfigFile)
	if errReadEnv != nil {
		return ""
	}
	strVal, ok := envMap[key]
	if !ok {
		return ""
	}
	return strVal
}

func EnvAsInt(key string, defaultVal int) int {
	strVal := Env(key)
	if val, err := strconv.Atoi(strVal); err == nil {
		return val
	}
	return defaultVal
}

func EnvAsDuration(key string, defaultVal time.Duration) time.Duration {
	strVal := Env(key)
	if val, err := time.ParseDuration(strVal); err == nil {
		return val
	}
	return defaultVal
}
