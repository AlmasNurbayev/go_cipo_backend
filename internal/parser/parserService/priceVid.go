package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) PriceVidService(mainStruct *xmltypes.OfferType,
	registrator_id int64) error {

	op := "PriceVidService"
	log := s.log.With(slog.String("op", op))

	NewPriceVids := partParsers.PriceVidParser(mainStruct, registrator_id)

	// берем из базы имеющие записи и проверяем на дубликаты
	existsPriceVids, err := s.storage.ListPriceVid(s.ctx)
	if err != nil {
		log.Error("Error selecting price_vids:", slog.String("err", err.Error()))
		return err
	}
	log.Info("exist price_vid: " + strconv.Itoa(len(existsPriceVids)))

	existsMap := make(map[string]models.PriceVidEntity)
	for _, e := range existsPriceVids {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.PriceVidEntity

	for _, n := range NewPriceVids {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdatePriceVid(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := s.storage.CreatePriceVid(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
			}
		}
	}

	log.Info("priceVid parsing: ", slog.Int("count", len(NewPriceVids)))
	log.Info("== Duplicated and updated priceVids: ", slog.Int("count", len(toUpdate)))
	log.Info("== Created new priceVids: ", slog.Int("count", len(toCreate)))

	return nil
}
