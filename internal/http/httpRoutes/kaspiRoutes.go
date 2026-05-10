package httproutes

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/http/middleware"
	"github.com/gofiber/fiber/v3"
)

type handlersKaspiProvider interface {
	KaspiAddCategory(c fiber.Ctx) error
	KaspiListCategory(c fiber.Ctx) error
	KaspiUpdateCategory(c fiber.Ctx) error

	KaspiAddOrganization(c fiber.Ctx) error
	KaspiListOrganization(c fiber.Ctx) error

	ListKaspiProducts(c fiber.Ctx) error
	KaspiGetByIdCategory(c fiber.Ctx) error

	KaspiExportProducts(c fiber.Ctx) error
}

func RegisterKaspiRoutes(app *fiber.App, handler handlersKaspiProvider, log *slog.Logger) {
	api := app.Group("/api")
	api.Post("/kaspi/category/", middleware.RequireAuth, handler.KaspiAddCategory)
	log.Info("POST /api/kaspi/category/")

	api.Get("/kaspi/categories/", middleware.RequireAuth, handler.KaspiListCategory)
	log.Info("GET /api/kaspi/categories/")

	api.Get("/kaspi/category/:id", middleware.RequireAuth, handler.KaspiGetByIdCategory)
	log.Info("GET /api/kaspi/category/:id")

	api.Put("/kaspi/category/", middleware.RequireAuth, handler.KaspiUpdateCategory)
	log.Info("PUT /api/kaspi/category/")

	api.Post("/kaspi/organization/", middleware.RequireAuth, handler.KaspiAddOrganization)
	log.Info("POST /api/kaspi/organization/")

	api.Get("/kaspi/organization/", middleware.RequireAuth, handler.KaspiListOrganization)
	log.Info("GET /api/kaspi/organization/")

	api.Get("/kaspi/products/", middleware.RequireAuth, handler.ListKaspiProducts)
	log.Info("GET /api/kaspi/products/")

	api.Post("/kaspi/export-products/", middleware.RequireAuth, handler.KaspiExportProducts)
	log.Info("POST /api/kaspi/export-products/")
}
