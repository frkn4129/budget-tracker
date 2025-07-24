package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string        // ":3000"
	JWTSecret   string        // JWT için gizli anahtar
	DBHost      string        // "localhost"
	DBPort      int           // 5432
	DBUser      string        // "postgres"
	DBPassword  string        // "postgres"
	DBName      string        // "elektriklihayatlar"
	DBSSLMode   string        // "disable"
	TokenExpiry time.Duration // 72 * time.Hour
}

// LoadConfig loads the environment variables and returns a Config object
func LoadConfig() (*Config, error) {
	// .env dosyasını yükle
	err := godotenv.Load("/Users/gurkanseker/Desktop/budget-tracker/auth-service/.env")
	if err != nil {
		return nil, err
	}

	// Port'u int olarak oku
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432 // default
	}

	// Token süresi dakika cinsinden okunur
	expiryMinutes, err := strconv.Atoi(os.Getenv("JWT_EXP_MINUTES"))
	if err != nil {
		expiryMinutes = 4320 // default: 72 saat
	}

	cfg := &Config{
		AppPort:     os.Getenv("APP_PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      dbPort,
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DBSSLMode:   os.Getenv("DB_SSLMODE"),
		TokenExpiry: time.Duration(expiryMinutes) * time.Minute,
	}

	if cfg.JWTSecret == "" {
		return nil, &ConfigError{Message: "JWT_SECRET is not set in environment"}
	}

	return cfg, nil
}

type ConfigError struct {
	Message string
}

func (e *ConfigError) Error() string {
	return e.Message
}
