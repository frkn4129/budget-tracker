package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/discovery-service/internal"
)

func main() {
	app := fiber.New()

	app.Get("/discover/:name", internal.DiscoverService)

	app.Listen(":4001")
}
