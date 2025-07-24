package routes

import (
	"gateway-service/internal/config"
	"gateway-service/internal/discovery"
	"gateway-service/internal/middleware"
	"gateway-service/internal/proxy"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// create resolvers that query discovery-service each time (simple implementation)
	authResolver := func() (string, error) {
		return discovery.GetServiceAddress(cfg.DiscoveryServiceURL, "auth-service")
	}
	budgetResolver := func() (string, error) {
		return discovery.GetServiceAddress(cfg.DiscoveryServiceURL, "budget-service")
	}

	// 🔓 Public Routes (Auth Servisi)
	app.Post("/api/auth/register", proxy.DynamicForward(authResolver))
	app.Post("/api/auth/login", proxy.DynamicForward(authResolver))

	// 🔐 Protected Routes
	authorized := app.Group("/api", middleware.JWTMiddleware(cfg.JWTSecret))

	authorized.Get("/auth/me", proxy.DynamicForward(authResolver)) // JWT korumalı

	// Budget Service: forward all under /expenses, including root
	authorized.All("/expenses*", proxy.DynamicForward(budgetResolver))
}
