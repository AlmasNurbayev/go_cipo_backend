package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) GetProductsOnlyQnt(ctx context.Context) (dto.ProductsOnlyQntResponse, error) {
	op := "Service.GetProductsOnlyQnt"
	log := s.log.With("op", op)

	res := dto.ProductsOnlyQntResponse{}

	lastOfferRegistrator, err := s.postgresStorage.GetLastOfferRegistrator(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return res, err
	}

	data, err := s.qntPriceStorage.ListProductsOnlyQnt(ctx, lastOfferRegistrator.Id)
	if err != nil {
		log.Error("Error", "err", err.Error())
		return dto.ProductsOnlyQntResponse{}, err
	}

	err = copier.Copy(&res.Data, &data)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return res, err
	}

	res.Count = len(res.Data)

	return res, nil
}
