package httphandlers

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/validate"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetProductBy(c fiber.Ctx) error {
	op := "HttpHandlers.GetProductBy"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateParams(c, &dto.ProductByIdQueryRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	idString := c.Query("id")
	if idString != "" {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			log.Warn(err.Error())
			return c.Status(400).SendString(errorsShare.ErrBadRequest.Message)
		}

		res, err := h.service.GetProductById(c, id)
		if err != nil {
			if err == errorsShare.ErrProductNotFound.Error {
				return c.Status(404).SendString(errorsShare.ErrProductNotFound.Message + " id " + idString)
			}
			log.Error("", slog.String("err", err.Error()))
			return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
		}
		return c.Status(200).JSON(res)
	}

	name1CString := c.Query("name_1c")
	if name1CString != "" {

		res, err := h.service.GetProductByName1c(c, name1CString)
		if err != nil {
			if err == errorsShare.ErrProductNotFound.Error {
				return c.Status(404).SendString(errorsShare.ErrProductNotFound.Message + " name_1c " + name1CString)
			}
			log.Error("", slog.String("err", err.Error()))
			return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
		}
		return c.Status(200).JSON(res)
	}

	return c.Status(400).SendString(errorsShare.ErrBadRequest.Message)

}
