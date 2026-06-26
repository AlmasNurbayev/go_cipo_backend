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

	lastOfferRegistrator, err := s.postgresStorage.GetLastOfferRegistrator(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	qntPriceRegistry, err := s.qntPriceStorage.GetQntPriceRegistryByProductId(ctx, id, lastOfferRegistrator.Id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	qntPriceRegistryGroup, err := s.qntPriceStorage.GetQntPriceRegistryGroupByProductId(ctx, id, lastOfferRegistrator.Id)
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

	err = copier.Copy(&productByIdDto.Qnt_price_registry_group, &qntPriceRegistryGroup)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByIdDto, err
	}

	productByIdDto.Vid_modeli = dto.IdName1c{
		Id:      int(productByIdEntity.Vid_modeli.Id.ValueOrZero()),
		Name_1c: productByIdEntity.Vid_modeli.Name_1c.ValueOrZero(),
	}
	productByIdDto.Product_group = dto.IdName1c{
		Id:      int(productByIdEntity.Product_group.Id.ValueOrZero()),
		Name_1c: productByIdEntity.Product_group.Name_1c.ValueOrZero(),
	}

	return productByIdDto, nil
}

func (s *Service) GetProductByName1c(ctx context.Context, name_1c string) (dto.ProductByIdResponse, error) {
	op := "services.GetProductByName1c"
	log := s.log.With(slog.String("op", op))

	productByNameDto := dto.ProductByIdResponse{}

	productByIdEntity, err := s.postgresStorage.GetProductByName1c(ctx, name_1c)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	imagesEntity, err := s.postgresStorage.GetImagesByProductId(ctx, productByIdEntity.Id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	lastOfferRegistrator, err := s.postgresStorage.GetLastOfferRegistrator(ctx)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	qntPriceRegistry, err := s.qntPriceStorage.GetQntPriceRegistryByProductId(ctx, productByIdEntity.Id, lastOfferRegistrator.Id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	qntPriceRegistryGroup, err := s.qntPriceStorage.GetQntPriceRegistryGroupByProductId(ctx, productByIdEntity.Id, lastOfferRegistrator.Id)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	err = copier.Copy(&productByNameDto, &productByIdEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}
	err = copier.Copy(&productByNameDto.Image_registry, &imagesEntity)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	err = copier.Copy(&productByNameDto.Qnt_price_registry, &qntPriceRegistry)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	err = copier.Copy(&productByNameDto.Qnt_price_registry_group, &qntPriceRegistryGroup)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return productByNameDto, err
	}

	productByNameDto.Vid_modeli = dto.IdName1c{
		Id:      int(productByIdEntity.Vid_modeli.Id.ValueOrZero()),
		Name_1c: productByIdEntity.Vid_modeli.Name_1c.ValueOrZero(),
	}
	productByNameDto.Product_group = dto.IdName1c{
		Id:      int(productByIdEntity.Product_group.Id.ValueOrZero()),
		Name_1c: productByIdEntity.Product_group.Name_1c.ValueOrZero(),
	}

	return productByNameDto, nil
}
