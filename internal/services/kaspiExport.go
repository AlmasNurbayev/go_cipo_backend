package services

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/clients"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/guregu/null/v5"
)

func (s *Service) KaspiExportProducts(ctx context.Context, req dto.KaspiExportProductRequest) (dto.KaspiExportProductResponse, error) {
	op := "Service.KaspiExportProducts"
	log := s.log.With(slog.String("op", op))

	var response dto.KaspiExportProductResponse

	if len(req.Data) == 0 || len(req.ProductIds) == 0 {
		log.Warn("request body has no data or productids")
		return response, errors.New("request body has no data or productids")
	}

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
		return dto.KaspiExportProductResponse{}, err
	}

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		log.Error("Error unmarshalling response:", slog.String("err", err.Error()))
		return dto.KaspiExportProductResponse{}, err
	}

	log.Debug("Kaspi Export Products response:", slog.Any("response", response))
	// собираем ошибки в массив строк
	var ErrorsFrom []string
	if len(response.Errors) > 0 {
		for _, errorItem := range response.Errors {
			errItemByte, err := json.Marshal(errorItem)
			if err == nil {
				ErrorsFrom = append(ErrorsFrom, string(errItemByte))
			} else {
				log.Error("Error marshalling error item:", slog.String("err", err.Error()))
			}
		}
	}

	sendedBody, err := json.Marshal(req)
	if err != nil {
		log.Error("Error marshalling request:", slog.String("err", err.Error()))
		return response, err
	}

	newRegistry := models.KaspiExportGoodsRegistryEntity{
		ProductId:           req.ProductIds[0],
		KaspiOrganizationId: req.OrganizationId,
		SendedBody:          string(sendedBody),
		SendedCategory:      req.Data[0].Category,
		SendedStatus:        statusCode,
		ResponseId:          null.StringFrom(response.Code),
		ResponseStatus:      null.StringFrom(response.Status),
		Errors:              ErrorsFrom,
	}

	_, err = s.kaspiStorage.CreateKaspiExportGoodsRegistry(ctx, newRegistry)
	if err != nil {
		log.Error("Error creating registry:", slog.String("err", err.Error()))
		//return response, errorsShare.ErrInternalError.Error
	}

	return response, nil
}
