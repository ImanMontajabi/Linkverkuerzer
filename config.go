package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port         string
	DatabaseURL  string
	BaseURL      string
	Environment  string
	ShortCodeLen int
	MaxURLLength int
}

func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "3000"),
		DatabaseURL:  getEnv("DATABASE_URL", "urls.db"),
		BaseURL:      getEnv("BASE_URL", "http://localhost:3000"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		ShortCodeLen: getEnvAsInt("SHORT_CODE_LENGTH", 6),
		MaxURLLength: getEnvAsInt("MAX_URL_LENGTH", 2048),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if valueStr := os.Getenv(key); valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
