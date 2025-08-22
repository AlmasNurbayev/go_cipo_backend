package httphandlers

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetProductFilters(c fiber.Ctx) error {
	op := "HttpHandlers.GetUser"
	log := h.log.With(slog.String("op", op))

	res, err := h.service.GetProductFilters(c)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}
	return c.Status(200).JSON(res)

}
