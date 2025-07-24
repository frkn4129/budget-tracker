package main

import (
	"auth-service/application"
	"auth-service/infrastructure/postgres"
	"auth-service/interfaces/http"
	"auth-service/interfaces/http/routes"
	"auth-service/internal/config"
	"auth-service/internal/consul"
	"auth-service/internal/jwt"
	"auth-service/internal/logger"
	"log"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func main() {
	// 1. Config yüklenir (.env üzerinden)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err) // env okunmazsa logger yoktur, panic ver
	}

	// 2. Logger başlat
	log := logger.NewLogger()

	// 3. DB bağlantısı
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("❌ Failed to connect to DB", zap.Error(err))
	}

	// 4. Repo tanımı
	userRepo := postgres.NewPostgresUserRepository(db)

	// 5. JWT setup
	jwtMaker := jwt.NewJWTMaker(cfg.JWTSecret)

	// 6. Handler'lar
	registerHandler := application.NewRegisterHandler(userRepo)
	loginHandler := application.NewLoginHandler(userRepo, jwtMaker)

	// 7. Route tanımı
	authRoutes := routes.NewAuthRoutes(registerHandler, loginHandler)

	// 8. Fiber server başlat
	app := http.InitServer(authRoutes, jwtMaker)

	// 9. Consul register
	portStr := strings.TrimPrefix(cfg.AppPort, ":")
	port, _ := strconv.Atoi(portStr)
	go consul.RegisterToConsul("auth-service-"+portStr, "auth-service", "localhost", port)

	log.Info("🚀 Auth Service running", zap.String("port", cfg.AppPort))
	if err := app.Listen(cfg.AppPort); err != nil {
		log.Fatal("🔥 Server crashed", zap.Error(err))
	}
}
