package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) GetNewsById(ctx context.Context, id int64) (dto.NewsIDItemResponse, error) {
	op := "services.GetNewsById"
	log := s.log.With(slog.String("op", op))

	dto := dto.NewsIDItemResponse{}

	entity, err := s.postgresStorage.GetNewsById(ctx, id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return dto, err
	}

	err = copier.Copy(&dto, &entity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return dto, err
	}

	return dto, nil
}

func (s *Service) ListNews(ctx context.Context, count int64) ([]dto.NewsItemResponse, error) {
	op := "services.ListNews"
	log := s.log.With(slog.String("op", op))

	dto := []dto.NewsItemResponse{}

	entity, err := s.postgresStorage.ListNews(ctx, count)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return dto, err
	}

	err = copier.Copy(&dto, &entity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return dto, err
	}

	return dto, nil
}
