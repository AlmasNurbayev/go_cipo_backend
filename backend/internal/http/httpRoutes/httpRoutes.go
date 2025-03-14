package httproutes

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type handlersProvider interface {
	GetUserById(c fiber.Ctx) error
	GetUserSearch(c fiber.Ctx) error

	GetProductFilters(c fiber.Ctx) error
	GetStores(c fiber.Ctx) error

	GetProductById(c fiber.Ctx) error
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

	log.Info("/api")
	api := app.Group("/api")
	log.Info("/user")
	api.Get("/user/search/", handler.GetUserSearch)

	log.Info("/user/:id?")
	api.Get("/user/:id?", handler.GetUserById)

	api.Get("/user/search/", handler.GetUserSearch)
	log.Info("/user/search/")

	api.Get("/productsFilter", handler.GetProductFilters)
	log.Info("/productsFilter")

	api.Get("/stores", handler.GetStores)
	log.Info("/stores")

	log.Info("/product/")
	api.Get("/product/", handler.GetProductById)

	log.Info("/productsNews/")
	api.Get("/productsNews/", handler.ListProductNews)

	log.Info("/newsID/")
	api.Get("/newsID/", handler.GetNewsById)

	log.Info("/news/")
	api.Get("/news/", handler.ListNews)

	log.Info("/products/")
	api.Get("/products/", handler.ListProducts)

}
