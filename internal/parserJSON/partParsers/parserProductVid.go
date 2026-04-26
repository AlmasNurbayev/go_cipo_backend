package partParsers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
)

type storageV interface {
	ListProductVid(context.Context) ([]models.ProductVidEntity, error)
	CreateProductVid(context.Context, models.ProductVidEntity) (int64, error)
	UpdateProductVid(context.Context, models.ProductVidEntity) error
}

func ParserProductVids(Log *slog.Logger, ctx context.Context, storage storageV, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.ParserProductVids"
	Log = Log.With(slog.String("op", op))

	NewProductVids := UniqueProductVids(data, registrator_id)
	// берем из базы имеющие записи и проверяем на дубликаты
	existsProductVids, err := storage.ListProductVid(ctx)
	if err != nil {
		Log.Error("Error selecting product_vids:", slog.String("err", err.Error()))
		return err
	}
	Log.Debug("exist product_vids: " + strconv.Itoa(len(existsProductVids)))

	existsMap := make(map[string]models.ProductVidEntity)
	for _, e := range existsProductVids {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.ProductVidEntity

	for _, n := range NewProductVids {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateProductVid(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := storage.CreateProductVid(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		}
	}

	Log.Debug("ProductVid parsing: ", slog.Int("count", len(NewProductVids)))
	Log.Debug("Duplicated and updated ProductVid: ", slog.Int("count", len(toUpdate)))
	Log.Info("Created new ProductVid: ", slog.Int("count", len(toCreate)))

	return nil
}

func UniqueProductVids(data dto.StockJSON,
	registrator_id int64) []models.ProductVidEntity {

	productVids := make([]models.ProductVidEntity, 0)

	root := data.Qnt
	tempMap := make(map[string]string)

	for _, item := range root {
		if item.VidGUID == "" {
			continue
		}
		if _, exists := tempMap[item.VidGUID]; !exists {
			tempMap[item.VidGUID] = item.VidGUID
			productVids = append(productVids, models.ProductVidEntity{
				Id_1c:          item.VidGUID,
				Name_1c:        item.VidName,
				Registrator_id: registrator_id,
			})
		}
	}
	return productVids
}
