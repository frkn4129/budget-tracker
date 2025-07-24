package middleware

import (
	"auth-service/internal/jwt"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(jwtMaker *jwt.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		// "Bearer <token>" şeklindeyse ayır
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		claims, err := jwtMaker.VerifyToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// UserID'yi context'e ekle
		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}
