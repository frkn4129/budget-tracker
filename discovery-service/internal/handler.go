package internal

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ServiceInfo struct {
	Address        string `json:"Address"`
	ServiceAddress string `json:"ServiceAddress"`
	Port           int    `json:"ServicePort"`
}

func DiscoverService(c *fiber.Ctx) error {
	name := c.Params("name")
	services, err := GetServicesFromConsul(name)
	if err != nil || len(services) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("no service found for %s", name),
		})
	}

	addr := services[0].Address
	if addr == "" {
		addr = services[0].ServiceAddress
	}

	target := fmt.Sprintf("http://%s:%d", addr, services[0].Port)
	return c.JSON(fiber.Map{
		"service": name,
		"target":  target,
	})
}
