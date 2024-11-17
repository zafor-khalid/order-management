package config

import (
	"order-management/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env          string
	Port         string
	
	// Database configs
	DBHost       string
	DBPort       string
	DBName       string
	DBUser       string
	DBPassword   string
	
	// Auth configs
	JWTSecret    string
	JWTExpiresIn string
	
	// API configs
	APIKey       string
}

func LoadConfig() (*Config, error) {
	// Load .env file from project root
	if err := godotenv.Load(); err != nil {
		logger.Info("Error loading .env file: " + err.Error())
	}

	return &Config{
		Env:          getEnv("ENV", "development"),
		Port:         getEnv("PORT", "8080"),
		
		// Database configs
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBName:       getEnv("DB_NAME", "mydatabase"),
		DBUser:       getEnv("DB_USER", "myuser"),
		DBPassword:   getEnv("DB_PASSWORD", "mypassword"),
		
		// Auth configs
		JWTSecret:    getEnv("JWT_SECRET", ""),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "3600"),
		
		// API configs
		APIKey:       getEnv("API_KEY", ""),
	}, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
