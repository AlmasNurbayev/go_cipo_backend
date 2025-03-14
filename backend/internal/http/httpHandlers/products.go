package httphandlers

import (
	"log/slog"
	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/validate"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) ListProducts(c fiber.Ctx) error {
	op := "HttpHandlers.ListProductNews"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateQueryParams(c, &dto.ProductsQueryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	params := dto.ProductsQueryRequest{}
	queryMap := c.Queries()

	params.Take = utils.String2Int(queryMap["take"])
	params.Skip = utils.String2Int(queryMap["skip"])
	params.MinPrice = utils.String2Float32(queryMap["minPrice"])
	params.MaxPrice = utils.String2Float32(queryMap["maxPrice"])
	params.Product_group = utils.String2ArrayInt64(queryMap["product_group"], ",")
	params.Vid_modeli = utils.String2ArrayInt64(queryMap["vid_modeli"], ",")
	params.Size = utils.String2ArrayInt64(queryMap["size"], ",")
	params.Search_name = queryMap["search_name"]
	params.Sort = queryMap["sort"]

	if params.Sort != "" && !strings.Contains(params.Sort, "-") {
		log.Error(errorsShare.ErrMaxPriceLessMinPrice.Message)
		return c.Status(errorsShare.ErrSortBadFormat.Code).SendString(errorsShare.ErrSortBadFormat.Message)
	}

	if (params.MinPrice != 0) && (params.MaxPrice != 0) {
		if params.MaxPrice < params.MinPrice {
			log.Error(errorsShare.ErrMaxPriceLessMinPrice.Message)
			return c.Status(errorsShare.ErrMaxPriceLessMinPrice.Code).SendString(errorsShare.ErrMaxPriceLessMinPrice.Message)

		}
	}

	res, err := h.service.ListProducts(c.Context(), params)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(res)

}
