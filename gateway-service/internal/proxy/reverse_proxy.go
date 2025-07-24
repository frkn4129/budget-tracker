package proxy

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Forward(targetBaseURL string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Hedef URL'yi oluştur
		targetURL := targetBaseURL + c.OriginalURL()
		log.Printf("[proxy] Forward -> %s %s", c.Method(), targetURL)

		// Yeni HTTP isteği oluştur
		req, err := http.NewRequest(c.Method(), targetURL, strings.NewReader(string(c.Body())))
		if err != nil {
			log.Printf("[proxy] create request error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create request to target service",
			})
		}

		// Header'ları kopyala
		c.Request().Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})

		// HTTP client ile istek gönder
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("[proxy] forward error to %s: %v", targetURL, err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "Failed to reach target service",
			})
		}
		defer resp.Body.Close()

		// Cevap body'sini oku ve geri döndür
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[proxy] read response error from %s: %v", targetURL, err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read response from service",
			})
		}

		// Orijinal status code ve content-type ile döndür
		c.Set("Content-Type", resp.Header.Get("Content-Type"))
		return c.Status(resp.StatusCode).Send(body)
	}
}

// DynamicForward creates a proxy handler that resolves target base URL via resolver for each request.
func DynamicForward(resolver func() (string, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		baseURL, err := resolver()
		if err != nil || baseURL == "" {
			log.Printf("[proxy] resolver error: %v", err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "Failed to resolve target service",
			})
		}
		targetURL := baseURL + c.OriginalURL()
		log.Printf("[proxy] DynamicForward -> %s %s", c.Method(), targetURL)

		req, err := http.NewRequest(c.Method(), targetURL, strings.NewReader(string(c.Body())))
		if err != nil {
			log.Printf("[proxy] create request error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create request to target service",
			})
		}
		c.Request().Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("[proxy] forward error to %s: %v", targetURL, err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "Failed to reach target service",
			})
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[proxy] read response error from %s: %v", targetURL, err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to read response from service",
			})
		}

		c.Set("Content-Type", resp.Header.Get("Content-Type"))
		return c.Status(resp.StatusCode).Send(body)
	}
}
