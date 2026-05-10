package httphandlers

import (
	"encoding/json"
	"log/slog"
	"strconv"
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

	response, err := h.service.KaspiAddCategory(c, body)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		if err == errorsShare.ErrKaspiCategoryDuplicate.Error {
			return c.Status(errorsShare.ErrKaspiCategoryDuplicate.Code).SendString(errorsShare.ErrKaspiCategoryDuplicate.Message)
		} else if err.Error() == "category code not found" {
			return c.Status(400).SendString("category code not found")
		} else {
			return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
		}
	}
	return c.Status(201).JSON(response)
}

func (h *Handler) KaspiListCategory(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiListCategory"
	log := h.log.With(slog.String("op", op))

	response, err := h.service.KaspiListCategory(c)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) KaspiGetByIdCategory(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiGetByIdCategory"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateParams(c, &dto.ProductByIdQueryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	idString := c.Params("id")
	var id int64
	if idString == "" {
		log.Warn(errorsShare.ErrBadRequest.Message)
		return c.Status(errorsShare.ErrBadRequest.Code).SendString(errorsShare.ErrBadRequest.Message)
	}
	id, err = strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(errorsShare.ErrBadRequest.Message)
	}

	response, err := h.service.KaspiGetByIdCategory(c, id)
	if err != nil {
		if err == errorsShare.ErrKaspiCategoryNotFound.Error {
			return c.Status(errorsShare.ErrKaspiCategoryNotFound.Code).SendString(errorsShare.ErrKaspiCategoryNotFound.Message)
		}
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

	response, err := h.service.KaspiAddOrganization(c, body)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) KaspiListOrganization(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiListOrganization"
	log := h.log.With(slog.String("op", op))

	response, err := h.service.KaspiListOrganization(c)
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

	params.Id = utils.String2Int64(queryMap["id"])
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

	res, err := h.service.ListKaspiProducts(c, params)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(res)

}

func (h *Handler) KaspiUpdateCategory(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiUpdateCategory"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateBody(c, &dto.KaspiUpdateCategoryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	body := dto.KaspiUpdateCategoryRequest{}
	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	response, err := h.service.KaspiUpdateCategory(c, body)
	if err != nil {
		if err == errorsShare.ErrKaspiCategoryNotFound.Error {
			return c.Status(errorsShare.ErrKaspiCategoryNotFound.Code).SendString(errorsShare.ErrKaspiCategoryNotFound.Message)
		}
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) KaspiExportProducts(c fiber.Ctx) error {
	op := "HttpHandlers.KaspiExportProducts"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateBody(c, &dto.KaspiExportProductRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	body := dto.KaspiExportProductRequest{}
	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}
	log.Debug("ExportProductRequest body:", slog.Any("body", body))

	//return c.Status(200).JSON(body)

	response, err := h.service.KaspiExportProducts(c, body)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}
