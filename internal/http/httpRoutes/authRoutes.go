package httproutes

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type handlersAuthProvider interface {
	Register(c fiber.Ctx) error
}

func RegisterAuthRoutes(app *fiber.App, handler handlersAuthProvider, log *slog.Logger) {
	api := app.Group("/api")
	api.Post("/auth/register/", handler.Register)
	log.Info("POST /api/auth/register/")
}
