package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) GetStores(ctx context.Context) (dto.StoresResponse, error) {
	op := "services.GetStores"
	log := s.log.With(slog.String("op", op))

	storesDTO := dto.StoresResponse{}
	storesEntity, err := s.classifierStorage.ListStore(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return storesDTO, err
	}

	err = copier.Copy(&storesDTO.Stores, &storesEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return storesDTO, err
	}

	return storesDTO, nil
}
