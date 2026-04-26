package partParsers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
)

type StorageS interface {
	ListStore(context.Context) ([]models.StoreEntity, error)
	CreateStore(context.Context, models.StoreEntity) (int64, error)
	UpdateStoreFrom1C(context.Context, models.StoreEntity) error
}

func ParserStore(Log *slog.Logger, ctx context.Context, storage StorageS, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.ParserStore"
	Log = Log.With(slog.String("op", op))

	NewStores := UniqueStore(data, registrator_id)
	// берем из базы имеющие записи и проверяем на дубликаты
	existsStores, err := storage.ListStore(ctx)
	if err != nil {
		Log.Error("Error selecting stores:", slog.String("err", err.Error()))
		return err
	}
	Log.Debug("exist stores: " + strconv.Itoa(len(existsStores)))

	existsMap := make(map[string]models.StoreEntity)
	for _, e := range existsStores {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.StoreEntity

	for _, n := range NewStores {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateStoreFrom1C(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := storage.CreateStore(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		}
	}

	Log.Debug("Store parsing: ", slog.Int("count", len(NewStores)))
	Log.Debug("Duplicated and updated Store: ", slog.Int("count", len(toUpdate)))
	Log.Info("Created new Store: ", slog.Int("count", len(toCreate)))

	return nil
}

func UniqueStore(data dto.StockJSON, registrator_id int64) []models.StoreEntity {

	stores := make([]models.StoreEntity, 0)

	root := data.Qnt
	tempMap := make(map[string]string)

	for _, qnt := range root {
		for _, item := range qnt.Stocks {
			if item.Quantity == 0 {
				continue
			}
			if _, exists := tempMap[item.StockGUID]; !exists {
				tempMap[item.StockGUID] = item.StockGUID
				stores = append(stores, models.StoreEntity{
					Id_1c:          item.StockGUID,
					Name_1c:        item.Stock,
					Registrator_id: registrator_id,
				})
			}
		}
	}
	return stores
}
