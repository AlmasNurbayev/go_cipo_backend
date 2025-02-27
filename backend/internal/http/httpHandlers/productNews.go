package httphandlers

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/validate"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) ListProductNews(c fiber.Ctx) error {
	op := "HttpHandlers.ListProductNews"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateParams(c, &dto.ProductNewsQueryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	newsString := c.Query("news")
	count, err := strconv.ParseInt(newsString, 10, 64)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(errorsShare.ErrBadRequest.Message)
	}

	res, err := h.service.ListProductNews(c.Context(), count)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
	}
	return c.Status(200).JSON(res)

}
