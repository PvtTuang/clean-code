package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	App        App
	PostgreSQL PostgreSQL
}

type App struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func LoadConfigs() Configs {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return Configs{
		App: App{
			Host: getEnv("APP_HOST", "127.0.0.1"),
			Port: getEnv("APP_PORT", "3000"),
		},
		PostgreSQL: PostgreSQL{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Database: getEnv("DB_NAME", "mydb"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
