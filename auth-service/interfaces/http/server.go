package http

import (
	"auth-service/interfaces/http/routes"
	"auth-service/internal/jwt"

	"github.com/gofiber/fiber/v2"
)

func InitServer(authRoutes *routes.AuthRoutes, jwtMaker *jwt.JWTService) *fiber.App {
	app := fiber.New()

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	auth := api.Group("/auth")

	authRoutes.RegisterRoutes(auth, jwtMaker)

	return app
}
