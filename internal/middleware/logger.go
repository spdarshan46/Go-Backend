package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"user-management-api/pkg/logger"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Log request
		logger.Log.Info("Request received",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		)

		err := c.Next()

		// Log response
		duration := time.Since(start)
		logger.Log.Info("Request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
			zap.Int("bytes", len(c.Response().Body())),
		)

		return err
	}
}