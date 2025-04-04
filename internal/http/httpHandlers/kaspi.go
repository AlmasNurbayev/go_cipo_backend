package httphandlers

import (
	"encoding/json"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
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
