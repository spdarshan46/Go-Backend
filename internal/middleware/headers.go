package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Headers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-API-Version", "v1")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Set("Content-Security-Policy", "default-src 'self'")

		return c.Next()
	}
}