package partParsers

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/gofiber/fiber/v3/log"
)

type storageVM interface {
	ListVidModeli(context.Context) ([]models.VidModeliEntity, error)
	CreateVidModeli(context.Context, models.VidModeliEntity) (int64, error)
	UpdateVidModeli(context.Context, models.VidModeliEntity) error
}

func ParserVidModeli(Log *slog.Logger, ctx context.Context, storage storageVM, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.ParserVidModeli"
	Log = Log.With(slog.String("op", op))

	NewVidModelis := UniqueVidModeli(data, registrator_id)
	// берем из базы имеющие записи и проверяем на дубликаты
	existsVidModelis, err := storage.ListVidModeli(ctx)
	if err != nil {
		Log.Error("Error selecting vid_modelis:", slog.String("err", err.Error()))
		return err
	}
	Log.Debug("exist vid_modelis: " + strconv.Itoa(len(existsVidModelis)))

	existsMap := make(map[string]models.VidModeliEntity)
	for _, e := range existsVidModelis {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.VidModeliEntity

	for _, n := range NewVidModelis {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateVidModeli(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := storage.CreateVidModeli(ctx, n)
			if err != nil {
				Log.Error(err.Error())
			}
		}
	}

	log.Debug("VidModeli parsing: ", slog.Int("count", len(NewVidModelis)))
	log.Debug("Duplicated and updated VidModeli: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new VidModeli: ", slog.Int("count", len(toCreate)))

	return nil
}

func UniqueVidModeli(data dto.StockJSON,
	registrator_id int64) []models.VidModeliEntity {

	vidModelis := make([]models.VidModeliEntity, 0)

	root := data.Qnt
	tempMap := make(map[string]string)

	for _, product := range root {
		for _, item := range product.AdditionalProperties {
			if !strings.Contains(item.NameProperty, "ВидМодели") {
				continue
			}
			if item.GUIDProperty == "" {
				continue
			}
			if _, exists := tempMap[item.GUIDProperty]; !exists {
				tempMap[item.GUIDProperty] = item.GUIDProperty
				vidModelis = append(vidModelis, models.VidModeliEntity{
					Id_1c:          item.GUIDProperty,
					Name_1c:        item.StrValueProperty,
					Registrator_id: registrator_id,
				})
			}
		}
	}
	return vidModelis
}
