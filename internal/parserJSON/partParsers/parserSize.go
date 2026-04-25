package partParsers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/gofiber/fiber/v3/log"
)

type storageSize interface {
	ListSize(context.Context) ([]models.SizeEntity, error)
	CreateSize(context.Context, models.SizeEntity) (int64, error)
	UpdateSize(context.Context, models.SizeEntity) error
}

func ParserSize(Log *slog.Logger, ctx context.Context, storage storageSize, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.ParserVidModeli"
	Log = Log.With(slog.String("op", op))

	NewSizes := UniqueSize(data, registrator_id)
	// берем из базы имеющие записи и проверяем на дубликаты
	existsSizes, err := storage.ListSize(ctx)
	if err != nil {
		Log.Error("Error selecting sizes:", slog.String("err", err.Error()))
		return err
	}
	Log.Debug("exist sizes: " + strconv.Itoa(len(existsSizes)))

	existsMap := make(map[string]models.SizeEntity)
	for _, e := range existsSizes {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.SizeEntity

	for _, n := range NewSizes {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateSize(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := storage.CreateSize(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		}
	}

	log.Debug("Sizes parsing: ", slog.Int("count", len(NewSizes)))
	log.Debug("Duplicated and updated Sizes: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new Sizes: ", slog.Int("count", len(toCreate)))

	return nil
}

func UniqueSize(data dto.StockJSON, registrator_id int64) []models.SizeEntity {

	sizes := make([]models.SizeEntity, 0)

	root := data.Qnt
	tempMap := make(map[string]string)

	for _, qnt := range root {
		for _, item := range qnt.Stocks {
			if item.Quantity == 0 {
				continue
			}
			if _, exists := tempMap[item.CharGUID]; !exists {
				tempMap[item.CharGUID] = item.CharGUID
				sizes = append(sizes, models.SizeEntity{
					Id_1c:          item.CharGUID,
					Name_1c:        item.Char,
					Registrator_id: registrator_id,
				})
			}
		}
	}
	return sizes
}
