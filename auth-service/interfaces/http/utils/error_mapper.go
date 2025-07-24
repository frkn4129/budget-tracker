package utils

import (
	"auth-service/application"
	"github.com/gofiber/fiber/v2"
)

func MapErrorToFiber(err error) (int, fiber.Map) {
	switch err {
	case application.ErrEmailAlreadyExists:
		return fiber.StatusConflict, fiber.Map{"error": "Email already registered"}
	case application.ErrInvalidInput:
		return fiber.StatusBadRequest, fiber.Map{"error": "Invalid input"}
	default:
		return fiber.StatusInternalServerError, fiber.Map{"error": "Internal server error"}
	}
}
