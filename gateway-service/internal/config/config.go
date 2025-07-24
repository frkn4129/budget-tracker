package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GatewayPort         string
	JWTSecret           string
	DiscoveryServiceURL string
}

func LoadConfig() *Config {
	// .env yüklenir
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using system environment variables")
	}

	cfg := &Config{
		GatewayPort:         getEnv("GATEWAY_PORT", "4000"),
		JWTSecret:           getEnv("JWT_SECRET", ""),
		DiscoveryServiceURL: getEnv("DISCOVERY_URL", "http://localhost:4001"),
	}

	return cfg
}

// Yardımcı fonksiyon
func getEnv(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
