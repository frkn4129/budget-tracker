package http

import (
	"budget-tracker/application"
	"budget-tracker/interfaces/routes"

	"github.com/gofiber/fiber/v2"
)

func NewFiberServer(createHandler *application.CreateExpenseHandler) *fiber.App {
	app := fiber.New()

	// health
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })

	// Global middleware’ler (loglama, CORS, auth) eklenebilir
	// app.Use(middleware.Logger())

	// Routes register
	routes.RegisterExpenseRoutes(app, createHandler)

	return app
}
