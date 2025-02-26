package services

import (
	"context"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/jinzhu/copier"
)

func (s *Service) GetProductById(ctx context.Context, id int64) (dto.ProductByIdResponse, error) {
	op := "services.GetProductById"
	log := s.log.With(slog.String("op", op))

	productByIdDto := dto.ProductByIdResponse{}

	productByIdEntity, err := s.postgresStorage.GetProductById(ctx, id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	imagesEntity, err := s.postgresStorage.GetImagesByProductId(ctx, id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	// productGroup, err := s.postgresStorage.GetProductGroupById(ctx, productEntity.Product_group_id)
	// if err != nil {
	// 	log.Error("", slog.String("err", err.Error()))
	// 	return productDto, err
	// }

	// if productEntity.Vid_modeli_id.Valid {
	// 	vidModeliId := null.IntFrom(productEntity.Vid_modeli_id.Int64).Int64
	// 	vidModeli, err := s.postgresStorage.GetVidModeliById(ctx, vidModeliId)
	// 	if err != nil {
	// 		log.Error("", slog.String("err", err.Error()))
	// 		return productDto, err
	// 	}
	// 	err = copier.Copy(&productDto.Vid_modeli, &vidModeli)
	// 	if err != nil {
	// 		log.Error("", slog.String("err", err.Error()))
	// 		return productDto, err
	// 	}
	// }

	lastOfferRegistrator, err := s.postgresStorage.GetLastOfferRegistrator(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	qntPriceRegistry, err := s.postgresStorage.GetQntPriceRegistryByProductId(ctx, id, lastOfferRegistrator.Id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	err = copier.Copy(&productByIdDto, &productByIdEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}
	err = copier.Copy(&productByIdDto.Image_registry, &imagesEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}
	// err = copier.Copy(&productDto.Product_group, &productGroup)
	// if err != nil {
	// 	log.Error("", slog.String("err", err.Error()))
	// 	return productDto, err
	// }
	err = copier.Copy(&productByIdDto.Qnt_price_registry, &qntPriceRegistry)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	// TODO - "qnt_price_registry_group" section

	return productByIdDto, nil
}
