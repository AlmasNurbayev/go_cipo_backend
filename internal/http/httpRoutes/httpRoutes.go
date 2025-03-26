package httproutes

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type handlersProvider interface {
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

func RegisterRoutes(app *fiber.App, handler handlersProvider, log *slog.Logger) {

	cp := "registerRoutes"
	log = log.With(slog.String("cp", cp))
	log.Info("Register routes:")

	app.Get("/metrics", handler.GetMetrics)
	log.Info("/metrics")

	log.Info("/api")
	api := app.Group("/api")
	log.Info("/api/user")
	api.Get("/user/search/", handler.GetUserSearch)

	log.Info("/api/user/:id?")
	api.Get("/user/:id?", handler.GetUserById)

	api.Get("/api/user/search/", handler.GetUserSearch)
	log.Info("/user/search/")

	api.Get("/api/productsFilter", handler.GetProductFilters)
	log.Info("/productsFilter")

	api.Get("/api/stores", handler.GetStores)
	log.Info("/stores")

	log.Info("/api/product/")
	api.Get("/product/", handler.GetProductBy)

	log.Info("/api/productsNews/")
	api.Get("/productsNews/", handler.ListProductNews)

	log.Info("/api/newsID/")
	api.Get("/newsID/", handler.GetNewsById)

	log.Info("/api/news/")
	api.Get("/news/", handler.ListNews)

	log.Info("/api/products/")
	api.Get("/products/", handler.ListProducts)

}
