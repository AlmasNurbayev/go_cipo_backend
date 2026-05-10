package services

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/clients"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
)

func (s *Service) KaspiExportProducts(ctx context.Context, req dto.ExportProductRequest) (any, error) {
	op := "Service.KaspiExportProducts"
	log := s.log.With(slog.String("op", op))

	var response any

	organizations, err := s.kaspiStorage.ListKaspiOrganization(ctx)
	if err != nil {
		return response, errorsShare.ErrInternalError.Error
	}

	var token string
	for _, organization := range organizations {
		if organization.Id == req.OrganizationId {
			token, err = utils.DecryptToken(s.cfg.Auth.SECRET_BYTE, organization.Kaspi_api_token_hash)
			if err != nil {
				return response, errorsShare.ErrInternalError.Error
			}
			log.Debug("selected organization", slog.String("organization_name", organization.Name))
		}
	}
	if token == "" {
		// если что-то не так с токеном
		return response, errorsShare.ErrInternalError.Error
	}

	statusCode, body, err := clients.KaspiExportProducts(s.cfg, log, token, req)
	if err != nil {
		log.Error("Error exporting products:", slog.String("err", err.Error()))
		return dto.ExportProductResponse{}, err
	}

	log.Debug("Kaspi Export Products response:", slog.Int("status_code", statusCode), slog.String("body", body))

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		log.Error("Error unmarshalling response:", slog.String("err", err.Error()))
		return dto.ExportProductResponse{}, err
	}

	return response, nil
}
