package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = getEnv("PORT", "8000")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
