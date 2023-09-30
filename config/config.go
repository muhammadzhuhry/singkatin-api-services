package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port               string
	Service            string
	BasicAuthUsername  string
	BasicAuthPassword  string
	MySQLConnectionURL string
	Timezone           string
	Domain             string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port = getEnv("PORT", "8000")
	Service = getEnv("SERVICE", "myTodoService")
	BasicAuthUsername = getEnv("BASIC_AUTH_USERNAME", "")
	BasicAuthPassword = getEnv("BASIC_AUTH_PASSWORD", "")
	MySQLConnectionURL = getEnv("MYSQL_CONNECTION_URL", "")
	Timezone = getEnv("TIMEZONE", "Asia/Bangkok")
	Domain = getEnv("DOMAIN", "localhost:"+Port+"/")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
