package partParsers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/gofiber/fiber/v3/log"
)

type storageC interface {
	ListProductGroup(context.Context) ([]models.ProductsGroupEntity, error)
	CreateProductGroup(context.Context, models.ProductsGroupEntity) (int64, error)
	UpdateProductGroup(context.Context, models.ProductsGroupEntity) error
}

func ParserProductGroups(Log *slog.Logger, ctx context.Context, storage storageC, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.ParserProductGroups"
	Log = Log.With(slog.String("op", op))

	NewProductGroups := UniqueProductGroups(data, registrator_id)
	// берем из базы имеющие записи и проверяем на дубликаты
	existsProductGroups, err := storage.ListProductGroup(ctx)
	if err != nil {
		log.Error("Error selecting product_groups:", slog.String("err", err.Error()))
		return err
	}
	log.Debug("exist product_groups: " + strconv.Itoa(len(existsProductGroups)))

	existsMap := make(map[string]models.ProductsGroupEntity)
	for _, e := range existsProductGroups {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.ProductsGroupEntity

	for _, n := range NewProductGroups {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateProductGroup(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := storage.CreateProductGroup(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		}
	}

	log.Debug("ProductGroup parsing: ", slog.Int("count", len(NewProductGroups)))
	log.Debug("Duplicated and updated ProductGroup: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new ProductGroup: ", slog.Int("count", len(toCreate)))

	return nil
}

func UniqueProductGroups(data dto.StockJSON,
	registrator_id int64) []models.ProductsGroupEntity {

	productGroups := make([]models.ProductsGroupEntity, 0)

	root := data.Qnt
	tempMap := make(map[string]string)

	for _, item := range root {
		if item.TovarGroupGUID == "" {
			continue
		}
		if _, exists := tempMap[item.TovarGroupGUID]; !exists {
			tempMap[item.TovarGroupGUID] = item.TovarGroupGUID
			productGroups = append(productGroups, models.ProductsGroupEntity{
				Id_1c:          item.TovarGroupGUID,
				Name_1c:        item.TovarGroupName,
				Registrator_id: registrator_id,
			})
		}
	}
	return productGroups
}
