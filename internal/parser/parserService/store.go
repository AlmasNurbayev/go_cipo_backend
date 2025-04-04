package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) StoreService(mainStruct *xmltypes.OfferType,
	registrator_id int64) error {

	op := "StoreService"
	log := s.log.With(slog.String("op", op))

	NewStores := partParsers.StoreParser(mainStruct, registrator_id)

	// берем из базы имеющие записи и проверяем на дубликаты
	existsStores, err := s.storage.ListStore(s.ctx)
	if err != nil {
		s.log.Error("Error selecting stores:", slog.String("err", err.Error()))
		return err
	}
	s.log.Debug("exist stores: " + strconv.Itoa(len(existsStores)))

	existsMap := make(map[string]models.StoreEntity)
	for _, e := range existsStores {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.StoreEntity

	for _, n := range NewStores {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateStoreFrom1C(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateStore(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		}
	}

	log.Debug("Store parsing: ", slog.Int("count", len(NewStores)))
	log.Debug("Duplicated and updated stores: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new stores: ", slog.Int("count", len(toCreate)))

	return nil
}
