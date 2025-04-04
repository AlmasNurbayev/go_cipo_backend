package httproutes

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type handlersKaspiProvider interface {
	KaspiAddCategory(c fiber.Ctx) error
	KaspiListCategory(c fiber.Ctx) error
	KaspiAddOrganization(c fiber.Ctx) error
	KaspiListOrganization(c fiber.Ctx) error
}

func RegisterKaspiRoutes(app *fiber.App, handler handlersKaspiProvider, log *slog.Logger) {
	api := app.Group("/api")
	api.Post("/kaspi/category/", handler.KaspiAddCategory)
	log.Info("POST /api/kaspi/category/")

	api.Get("/kaspi/category/", handler.KaspiListCategory)
	log.Info("GET /api/kaspi/category/")

	api.Post("/kaspi/organization/", handler.KaspiAddOrganization)
	log.Info("POST /api/kaspi/organization/")

	api.Get("/kaspi/organization/", handler.KaspiListOrganization)
	log.Info("GET /api/kaspi/organization/")
}
