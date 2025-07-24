package main

import (
	"gateway-service/internal/config"
	"gateway-service/internal/consul"
	"gateway-service/routes"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Config yükle
	cfg := config.LoadConfig()

	// 2. Fiber başlat
	app := fiber.New()

	// health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })

	// 3. Route'ları tanımla
	routes.SetupRoutes(app, cfg)

	// 4. Consul register (async)
	port, _ := strconv.Atoi(cfg.GatewayPort)
	go consul.RegisterToConsul("gateway-service-"+cfg.GatewayPort, "gateway-service", "localhost", port)

	// 5. Uygulama başlat
	log.Printf("🚀 API Gateway listening on port %s", cfg.GatewayPort)
	if err := app.Listen(":" + cfg.GatewayPort); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
