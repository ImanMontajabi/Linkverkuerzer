package main

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	BaseURL     string
	Environment string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "3000"),
		DatabaseURL: getEnv("DATABASE_URL", "url.db"),
		BaseURL:     getEnv("BASE_URL", "http://localhost:3000"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
