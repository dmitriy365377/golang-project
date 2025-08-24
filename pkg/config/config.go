package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AuthServicePort string
	ChatServicePort string
	DatabaseURL     string
	JWTSecret       string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		AuthServicePort: getEnv("AUTH_SERVICE_PORT", ":50051"),
		ChatServicePort: getEnv("CHAT_SERVICE_PORT", ":50052"),
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/chatdb"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
