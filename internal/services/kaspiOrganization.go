package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/jinzhu/copier"
)

func (s *Service) KaspiAddOrganization(ctx context.Context, data dto.KaspiAddOrganizationRequest) (dto.KaspiAddOrganizationResponse, error) {
	op := "services.KaspiAddOrganization"
	log := s.log.With(slog.String("op", op))

	var newOrg models.KaspiOrganizationEntity
	var result dto.KaspiAddOrganizationResponse

	newOrg.Name = data.Name
	newOrg.Kaspi_id = data.Kaspi_id
	token := utils.EncryptToken(s.cfg.Auth.SECRET_BYTE, data.Kaspi_api_token)
	newOrg.Kaspi_api_token_hash = token

	response, err := s.postgresStorage.CreateKaspiOrganization(ctx, newOrg)
	if err != nil {
		return result, errorsShare.ErrInternalError.Error
	}

	err = copier.Copy(&result, &response)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return result, err
	}

	return result, nil
}

func (s *Service) KaspiListOrganization(ctx context.Context) (dto.KaspiListOrganizationResponse, error) {
	op := "services.KaspiListOrganization"
	log := s.log.With(slog.String("op", op))

	var result dto.KaspiListOrganizationResponse

	response, err := s.postgresStorage.ListKaspiOrganization(ctx)
	if err != nil {
		return result, errorsShare.ErrInternalError.Error
	}

	err = copier.Copy(&result.Data, &response)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return result, err
	}

	return result, nil
}
