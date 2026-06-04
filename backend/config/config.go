package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBPath string
}

// Load lê o arquivo .env e joga os valores para dentro da struct Config.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado")
	}

	return Config{
		Port:   getEnvOrDefault("PORT", "8080"),
		DBPath: getEnvOrDefault("DB_PATH", "./database/figurinhasCopa.db"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}