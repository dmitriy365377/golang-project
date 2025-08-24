package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AuthServicePort string
	ChatServicePort string
	DatabaseURL     string
	JWTSecret       string
	CORSOrigins     []string
	CORSMethods     []string
	CORSHeaders     []string
	// Cookie настройки
	CookieSecure   bool
	CookieDomain   string
	CookieSameSite string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		AuthServicePort: getEnv("AUTH_SERVICE_PORT", ":50051"),
		ChatServicePort: getEnv("CHAT_SERVICE_PORT", ":50052"),
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/chatdb"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"),
		CORSOrigins:     getEnvSlice("CORS_ORIGINS", []string{"*"}),
		CORSMethods:     getEnvSlice("CORS_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		CORSHeaders:     getEnvSlice("CORS_HEADERS", []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}),
		CookieSecure:    getEnvBool("COOKIE_SECURE", false),
		CookieDomain:    getEnv("COOKIE_DOMAIN", "localhost"),
		CookieSameSite:  getEnv("COOKIE_SAME_SITE", "lax"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true"
	}
	return defaultValue
}
