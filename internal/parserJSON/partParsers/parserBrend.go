package partParsers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
)

type storageB interface {
	ListBrend(context.Context) ([]models.BrendEntity, error)
	CreateBrend(context.Context, models.BrendEntity) (int64, error)
	UpdateBrend(context.Context, models.BrendEntity) error
}

func ParserBrend(Log *slog.Logger, ctx context.Context, storage storageB, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.ParserBrend"
	Log = Log.With(slog.String("op", op))

	NewBrends := UniqueBrend(data, registrator_id)
	// берем из базы имеющие записи и проверяем на дубликаты
	existsBrends, err := storage.ListBrend(ctx)
	if err != nil {
		Log.Error("Error selecting brends:", slog.String("err", err.Error()))
		return err
	}
	Log.Debug("exist brends: " + strconv.Itoa(len(existsBrends)))

	existsMap := make(map[string]models.BrendEntity)
	for _, e := range existsBrends {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.BrendEntity

	for _, n := range NewBrends {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateBrend(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := storage.CreateBrend(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		}
	}

	Log.Debug("Brend parsing: ", slog.Int("count", len(NewBrends)))
	Log.Debug("Duplicated and updated Brends: ", slog.Int("count", len(toUpdate)))
	Log.Info("Created new Brends: ", slog.Int("count", len(toCreate)))

	return nil
}

func UniqueBrend(data dto.StockJSON, registrator_id int64) []models.BrendEntity {

	brends := make([]models.BrendEntity, 0)

	root := data.Qnt
	tempMap := make(map[string]string)

	for _, item := range root {
		if item.MarkaGUID == "" || item.MarkaName == "" {
			continue
		}
		if _, exists := tempMap[item.MarkaGUID]; !exists {
			tempMap[item.MarkaGUID] = item.MarkaGUID
			brends = append(brends, models.BrendEntity{
				Id_1c:          item.MarkaGUID,
				Name_1c:        item.MarkaName,
				Registrator_id: registrator_id,
			})

		}
	}
	return brends
}
