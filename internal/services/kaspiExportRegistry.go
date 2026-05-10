package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/jinzhu/copier"
)

func (s *Service) AddKaspiExportGoodsRegistry(ctx context.Context, data models.KaspiExportGoodsRegistryEntity) (int64, error) {
	op := "AddKaspiExportGoodsRegistry"
	log := s.log.With(slog.String("op", op))

	id, err := s.kaspiStorage.CreateKaspiExportGoodsRegistry(ctx, data)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return 0, err
	}

	return id, nil
}

func (s *Service) GetKaspiExportGoodsRegistryByProductId(ctx context.Context, productId int64) (dto.KaspiExportGoodsRegistryItem, error) {
	op := "GetKaspiExportGoodsRegistryByProductId"
	log := s.log.With(slog.String("op", op))

	var response dto.KaspiExportGoodsRegistryItem

	data, err := s.kaspiStorage.GetKaspiExportGoodsRegistryByProductId(ctx, productId)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return response, err
	}

	err = copier.Copy(&response, &data)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return response, err
	}

	return response, nil
}

func (s *Service) ListKaspiExportGoodsRegistry(ctx context.Context) ([]dto.KaspiExportGoodsRegistryItem, error) {
	op := "ListKaspiExportGoodsRegistry"
	log := s.log.With(slog.String("op", op))

	var response []dto.KaspiExportGoodsRegistryItem

	data, err := s.kaspiStorage.ListKaspiExportGoodsRegistry(ctx)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return response, err
	}

	err = copier.Copy(&response, &data)
	if err != nil {
		log.Error("error: ", slog.String("err", err.Error()))
		return response, err
	}

	return response, nil
}
