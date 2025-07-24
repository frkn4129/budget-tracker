package routes

import (
	"auth-service/application"
	"auth-service/interfaces/http/middleware"
	"auth-service/interfaces/http/utils"
	"auth-service/internal/jwt"

	"github.com/gofiber/fiber/v2"
)

type AuthRoutes struct {
	RegisterHandler *application.RegisterHandler
	LoginHandler    *application.LoginHandler
}

func NewAuthRoutes(register *application.RegisterHandler, login *application.LoginHandler) *AuthRoutes {
	return &AuthRoutes{
		RegisterHandler: register,
		LoginHandler:    login,
	}
}

func (r *AuthRoutes) RegisterRoutes(app fiber.Router, jwtMaker *jwt.JWTService) {
	app.Post("/register", r.handleRegister)
	app.Post("/login", r.handleLogin)
	// JWT korumalı endpoint
	app.Get("/me", middleware.JWTMiddleware(jwtMaker), r.handleMe)
	app.Get("/health", r.handleHealth)

}

func (r *AuthRoutes) handleHealth(c *fiber.Ctx) error {
	return c.SendString("ok")
}
func (r *AuthRoutes) handleRegister(c *fiber.Ctx) error {
	var cmd application.RegisterCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := r.RegisterHandler.Handle(&cmd)
	if err != nil {
		status, body := utils.MapErrorToFiber(err)
		return c.Status(status).JSON(body)
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (r *AuthRoutes) handleLogin(c *fiber.Ctx) error {
	var cmd application.LoginCommand
	if err := c.BodyParser(&cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := r.LoginHandler.Handle(&cmd)
	if err != nil {
		status, body := utils.MapErrorToFiber(err)
		return c.Status(status).JSON(body)
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (r *AuthRoutes) handleMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id") // middleware’de set edilmişti
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": userID,
	})
}
