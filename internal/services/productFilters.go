package services

import (
	"context"
	"log/slog"
	"sort"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/jinzhu/copier"
)

func (s *Service) GetProductFilters(ctx context.Context) (dto.ProductsFilterResponse, error) {
	op := "services.GetProductFilters"
	log := s.log.With(slog.String("op", op))

	productFilterDto := dto.ProductsFilterResponse{}

	sizeEntity, err := s.classifierStorage.ListSize(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}
	// фильтруем размеры если задана максимальная длина символов
	if s.cfg.HTTP.EXCLUDES_SIZES_LEN_MIN > 0 {
		filtered := make([]models.SizeEntity, 0)
		for _, v := range sizeEntity {
			if len(v.Name_1c) <= s.cfg.HTTP.EXCLUDES_SIZES_LEN_MIN {
				filtered = append(filtered, v)
			}
		}
		sizeEntity = filtered
	}

	sort.Slice(sizeEntity, func(i, j int) bool {
		return sizeEntity[i].Name_1c < sizeEntity[j].Name_1c
	})

	productGroupEntity, err := s.classifierStorage.ListProductGroup(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}
	sort.Slice(productGroupEntity, func(i, j int) bool {
		return productGroupEntity[i].Name_1c < productGroupEntity[j].Name_1c
	})

	storeEntity, err := s.classifierStorage.ListStore(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	vidModeliEntity, err := s.classifierStorage.ListVidModeli(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}
	sort.Slice(vidModeliEntity, func(i, j int) bool {
		return vidModeliEntity[i].Name_1c < vidModeliEntity[j].Name_1c
	})

	brendEntity, err := s.classifierStorage.ListBrend(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	nomVids, err := s.classifierStorage.ListProductNomvids(ctx)
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

	err = copier.Copy(&productFilterDto.Nom_vid, &nomVids)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productFilterDto, err
	}

	return productFilterDto, nil

}
