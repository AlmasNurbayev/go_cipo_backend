package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) GetProductFilters(ctx context.Context) (dto.ProductsFilterResponse, error) {
	op := "services.GetProductFilters"
	log := s.log.With(slog.String("op", op))

	productFilterDto := dto.ProductsFilterResponse{}

	sizeEntity, err := s.postgresStorage.ListSize(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	productGroupEntity, err := s.postgresStorage.ListProductGroup(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	storeEntity, err := s.postgresStorage.ListStore(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	vidModeliEntity, err := s.postgresStorage.ListVidModeli(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	brendEntity, err := s.postgresStorage.ListBrend(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	err = copier.Copy(&productFilterDto.Size, &sizeEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}
	err = copier.Copy(&productFilterDto.Product_group, &productGroupEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	err = copier.Copy(&productFilterDto.Store, &storeEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	err = copier.Copy(&productFilterDto.Brend, &brendEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	err = copier.Copy(&productFilterDto.Vid_modeli, &vidModeliEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	return productFilterDto, nil

}
