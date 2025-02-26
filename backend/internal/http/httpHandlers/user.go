package httphandlers

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/validate"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetUserById(c fiber.Ctx) error {
	op := "HttpHandlers.GetUser"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateParams(c, &dto.UserRequestParams{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	for i := 1; i <= 2000; i++ {
		//fmt.Printf("Step %d\n", i)  // Логирование шага
		//time.Sleep(1 * time.Millisecond) // Задержка 1 секунда
	}

	idString := c.Params("id")
	res := dto.UserResponse{}

	if idString != "" {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			log.Warn(err.Error())
			return c.Status(400).SendString(errorsShare.ErrBadRequest.Message)
		}

		res, err = h.service.GetUserByIdService(c.Context(), id)
		if err != nil {
			log.Warn(err.Error())
			if err == errorsShare.ErrUserNotFound.Error {
				return c.Status(404).SendString(errorsShare.ErrUserNotFound.Message)
			}
			return c.Status(500).SendString(err.Error())
		}

	}
	return c.Status(200).JSON(res)
}

func (h *Handler) GetUserSearch(c fiber.Ctx) error {
	op := "HttpHandlers.GetUserSearch"
	log := h.log.With(slog.String("op", op))

	err := validate.ValidateQueryParams(c, &dto.UserRequestQueryParams{})
	if err != nil {
		log.Warn(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	for i := 1; i <= 2000; i++ {
		//fmt.Printf("Step %d\n", i)  // Логирование шага
		//time.Sleep(1 * time.Millisecond) // Задержка 1 секунда
	}

	nameString := c.Query("name")
	res := dto.UserResponse{}

	if nameString != "" {
		res, err = h.service.GetUserByNameService(c.Context(), nameString)
		if err != nil {
			log.Warn(err.Error())
			if err == errorsShare.ErrUserNotFound.Error {
				return c.Status(404).SendString(errorsShare.ErrUserNotFound.Message)
			}
			return c.Status(500).SendString(err.Error())
		}
	}
	return c.Status(200).JSON(res)
}
