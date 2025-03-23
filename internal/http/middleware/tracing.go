package middleware

import (
	"context"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type traceIDKey string

func RequestTracingMiddleware(log *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		traceID := uuid.New().String()

		ctx := context.WithValue(c.Context(), traceIDKey("traceID"), traceID)
		c.Set("X-Trace-ID", traceID)
		start := time.Now()

		err := c.Next()

		statusCode := c.Response().StatusCode()
		duration := time.Since(start)

		if err != nil {
			log.ErrorContext(ctx, err.Error())
			return err
		}

		log.Info("imcoming request",
			"traceId", traceID,
			"ip", c.IP(),
			"method", c.Method(),
			"path", c.Path(),
			"originalUrl", c.OriginalURL(),
			"status", statusCode,
			"duration", duration)
		return nil
	}
}
