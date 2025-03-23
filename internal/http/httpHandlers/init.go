package httphandlers

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
)

type servicesProvider interface {
	GetUserByIdService(ctx context.Context, id int64) (dto.UserResponse, error)
	GetUserByNameService(ctx context.Context, name string) (dto.UserResponse, error)

	GetProductFilters(ctx context.Context) (dto.ProductsFilterResponse, error)
	GetStores(ctx context.Context) (dto.StoresResponse, error)

	GetProductById(ctx context.Context, id int64) (dto.ProductByIdResponse, error)
	GetProductByName1c(ctx context.Context, name_1c string) (dto.ProductByIdResponse, error)
	ListProductNews(ctx context.Context, count int64) ([]dto.ProductNewsResponse, error)

	GetNewsById(ctx context.Context, id int64) (dto.NewsIDItemResponse, error)
	ListNews(ctx context.Context, count int64) ([]dto.NewsItemResponse, error)

	ListProducts(ctx context.Context, params dto.ProductsQueryRequest) (dto.ProductsResponse, error)
}

type Handler struct {
	log     *slog.Logger
	service servicesProvider
}

func NewHandler(log *slog.Logger, service servicesProvider) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

// func (h *Handler) Tracing(c fiber.Ctx) error {
// 	statusCode := strconv.Itoa(c.Response().StatusCode())
// 	infoString := c.IP() + " " + c.Port() + " " + c.Method() + " " + statusCode + " " + c.Path()
// 	h.log.Info(infoString)
// 	return nil
// }
