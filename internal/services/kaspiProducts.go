package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) ListKaspiProducts(ctx context.Context, params dto.KaspiProductsQueryRequest) (dto.KaspiProductsResponse, error) {
	op := "services.ListKaspiProducts"
	log := s.log.With(slog.String("op", op))

	data := dto.KaspiProductsResponse{}

	registrator, err := s.postgresStorage.GetLastOfferRegistrator(ctx)
	if err != nil {
		log.Error("error", slog.String("err", err.Error()))
		return data, err
	}

	list, fullCount, err := s.postgresStorage.ListKaspiProductsSearch(ctx, registrator.Id, params)
	if err != nil {
		log.Error("error", slog.String("err", err.Error()))
		return data, err
	}

	err = copier.Copy(&data.Data, &list)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return data, err
	}
	data.Current_count = len(data.Data)
	data.Full_count = fullCount

	return data, nil

}
