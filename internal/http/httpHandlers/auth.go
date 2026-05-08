package httphandlers

import (
	"encoding/json"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/validate"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Register(c fiber.Ctx) error {
	op := "HttpHandlers.Register"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateBody(c, &dto.AuthRegisterRequest{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	body := dto.AuthRegisterRequest{}
	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	response, err := h.service.CreateUserService(c, body)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		if err == errorsShare.ErrUserAlreadyExists.Error {
			return c.Status(errorsShare.ErrUserAlreadyExists.Code).SendString(errorsShare.ErrUserAlreadyExists.Message)
		} else {
			return c.Status(500).SendString(errorsShare.ErrInternalError.Message)
		}
	}
	return c.Status(201).JSON(response)
}
