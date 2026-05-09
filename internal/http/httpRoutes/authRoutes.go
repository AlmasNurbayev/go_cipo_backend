package httproutes

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/http/middleware"
	"github.com/gofiber/fiber/v3"
)

type handlersAuthProvider interface {
	Register(c fiber.Ctx) error
	Login(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
}

func RegisterAuthRoutes(app *fiber.App, handler handlersAuthProvider, log *slog.Logger) {
	api := app.Group("/api")
	api.Post("/auth/register/", handler.Register)
	log.Info("POST /api/auth/register/")

	api.Post("/auth/login/", handler.Login)
	log.Info("POST /api/auth/login/")

	api.Post("/auth/logout/", middleware.RequireAuth, handler.Logout)
	log.Info("POST /api/auth/logout/")
}
