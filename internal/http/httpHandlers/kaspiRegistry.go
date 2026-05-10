package httphandlers

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) ListKaspiExportGoodsRegistry(c fiber.Ctx) error {
	op := "HttpHandlers.ListKaspiExportGoodsRegistry"
	log := h.log.With(slog.String("op", op))

	response, err := h.service.ListKaspiExportGoodsRegistry(c)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) GetKaspiExportGoodsRegistryByProductId(c fiber.Ctx) error {
	op := "HttpHandlers.GetKaspiExportGoodsRegistryByProductId"
	log := h.log.With(slog.String("op", op))

	productId, err := strconv.ParseInt(c.Params("product_id"), 10, 64)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}
	response, err := h.service.GetKaspiExportGoodsRegistryByProductId(c, productId)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}

	return c.Status(200).JSON(response)
}
