package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) ListProductNews(ctx context.Context, count int64) ([]dto.ProductNewsResponse, error) {
	op := "services.ListProductNews"
	log := s.log.With(slog.String("op", op))

	var data = []dto.ProductNewsResponse{}

	registrator, err := s.postgresStorage.GetLastOfferRegistrator(ctx)
	if err != nil {
		log.Error("error", slog.String("err", err.Error()))
		return data, err
	}

	list, err := s.postgresStorage.ListProductNews(ctx, registrator.Id, count)
	if err != nil {
		log.Error("error", slog.String("err", err.Error()))
		return data, err
	}

	err = copier.Copy(&data, &list)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return data, err
	}

	return data, nil
}
