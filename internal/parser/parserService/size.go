package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) SizeService(mainStruct *xmltypes.OfferType,
	registrator_id int64) error {

	op := "SizeService"
	log := s.log.With(slog.String("op", op))

	NewSizes := partParsers.SizeParser(mainStruct, registrator_id)

	// берем из базы имеющие записи и проверяем на дубликаты
	existsSizes, err := s.storage.ListSize(s.ctx)
	if err != nil {
		log.Error("Error selecting sizes:", slog.String("err", err.Error()))
		return err
	}
	log.Info("exist size: " + strconv.Itoa(len(existsSizes)))

	existsMap := make(map[string]models.SizeEntity)
	for _, e := range existsSizes {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.SizeEntity

	for _, n := range NewSizes {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateSize(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		} else {
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateSize(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}

	log.Info("size parsing: ", slog.Int("count", len(NewSizes)))
	log.Info("== Duplicated and updated sizes: ", slog.Int("count", len(toUpdate)))
	log.Info("== Created new sizes: ", slog.Int("count", len(toCreate)))

	return nil
}
