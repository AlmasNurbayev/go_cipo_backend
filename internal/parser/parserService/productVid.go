package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) ProductVidService(mainStruct *xmltypes.ImportType,
	registrator_id int64) error {

	op := "ProductVidService"
	log := s.log.With(slog.String("op", op))

	NewVids := partParsers.ProductVidsParser(mainStruct, registrator_id)

	// берем из базы имеющие записи и проверяем на дубликаты

	existsVids, err := s.storage.ListProductVid(s.ctx)
	if err != nil {
		log.Error("Error selecting product vids:", slog.String("err", err.Error()))
		return err
	}
	log.Debug("exist product vids: " + strconv.Itoa(len(existsVids)))

	existsMap := make(map[string]models.ProductVidEntity)
	for _, e := range existsVids {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.ProductVidEntity

	for _, n := range NewVids {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateProductVid(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateProductVid(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		}
	}

	log.Debug("ProductVid parsing: ", slog.Int("count", len(NewVids)))
	log.Debug("Duplicated and updated productVid: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new productVid: ", slog.Int("count", len(toCreate)))

	return nil
}
