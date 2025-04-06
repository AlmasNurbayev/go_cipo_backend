package httphandlers

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/validate"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) KaspiAddCategory(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiAddCategory"
	log := h.log.With(slog.String("op", op))
	err := validate.ValidateBody(c, &dto.KaspiAddCategoryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	body := dto.KaspiAddCategoryRequest{}
	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	response, err := h.service.KaspiAddCategory(c.Context(), body)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		if err == errorsShare.ErrKaspiCategoryDuplicate.Error {
			return c.Status(errorsShare.ErrKaspiCategoryDuplicate.Code).SendString(errorsShare.ErrKaspiCategoryDuplicate.Message)
		} else {
			return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
		}
	}
	return c.Status(200).JSON(response)
}

func (h *Handler) KaspiListCategory(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiListCategory"
	log := h.log.With(slog.String("op", op))

	response, err := h.service.KaspiListCategory(c.Context())
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) KaspiAddOrganization(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiAddOrganization"
	log := h.log.With(slog.String("op", op))
	err := validate.ValidateBody(c, &dto.KaspiAddOrganizationRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	body := dto.KaspiAddOrganizationRequest{}
	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	response, err := h.service.KaspiAddOrganization(c.Context(), body)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) KaspiListOrganization(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiListOrganization"
	log := h.log.With(slog.String("op", op))

	response, err := h.service.KaspiListOrganization(c.Context())
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) ListKaspiProducts(c fiber.Ctx) error {
	op := "HttpHandlers.ListKaspiProducts"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateQueryParams(c, &dto.ProductsQueryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	params := dto.KaspiProductsQueryRequest{}
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

	// проверяем формат сортировки
	if params.Sort != "" && !strings.Contains(params.Sort, "-") {
		log.Error(errorsShare.ErrMaxPriceLessMinPrice.Message)
		return c.Status(errorsShare.ErrSortBadFormat.Code).SendString(errorsShare.ErrSortBadFormat.Message)
	}

	// проверяем корректность Макс суммы
	if (params.MinPrice != 0) && (params.MaxPrice != 0) {
		if params.MaxPrice < params.MinPrice {
			log.Error(errorsShare.ErrMaxPriceLessMinPrice.Message)
			return c.Status(errorsShare.ErrMaxPriceLessMinPrice.Code).SendString(errorsShare.ErrMaxPriceLessMinPrice.Message)
		}
	}

	// проверяем наличие take, если нет то задает дефолтные 20
	if params.Take == 0 {
		params.Take = 20
	}

	res, err := h.service.ListKaspiProducts(c.Context(), params)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(res)

}
