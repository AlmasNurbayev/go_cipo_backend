package httproutes

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type handlersMainProvider interface {
	GetMetrics(c fiber.Ctx) error

	GetUserById(c fiber.Ctx) error
	GetUserSearch(c fiber.Ctx) error

	GetProductFilters(c fiber.Ctx) error
	GetStores(c fiber.Ctx) error

	GetProductBy(c fiber.Ctx) error
	ListProductNews(c fiber.Ctx) error

	GetNewsById(c fiber.Ctx) error
	ListNews(c fiber.Ctx) error

	ListProducts(c fiber.Ctx) error
	//Tracing(c fiber.Ctx) error
}

func RegisterMainRoutes(app *fiber.App, handler handlersMainProvider, log *slog.Logger) {

	cp := "registerRoutes"
	log = log.With(slog.String("cp", cp))
	log.Info("Register routes:")

	app.Get("GET /metrics", handler.GetMetrics)
	log.Info("/metrics")

	log.Info("/api")
	api := app.Group("/api")

	log.Info("GET /api/user")
	api.Get("/user/search/", handler.GetUserSearch)

	log.Info("GET /api/user/:id?")
	api.Get("/user/:id?", handler.GetUserById)

	log.Info("GET /api/user/search/")
	api.Get("/user/search/", handler.GetUserSearch)

	log.Info("GET /api/productsFilter")
	api.Get("/productsFilter", handler.GetProductFilters)

	log.Info("GET /api/stores")
	api.Get("/stores", handler.GetStores)

	log.Info("GET /api/product/")
	api.Get("/product/", handler.GetProductBy)

	log.Info("GET /api/productsNews/")
	api.Get("/productsNews/", handler.ListProductNews)

	log.Info("GET /api/newsID/")
	api.Get("/newsID/", handler.GetNewsById)

	log.Info("GET /api/news/")
	api.Get("/news/", handler.ListNews)

	log.Info("GET /api/products/")
	api.Get("/products/", handler.ListProducts)

}
