package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) VidModeliService(mainStruct *xmltypes.ImportType,
	registrator_id int64) error {

	op := "VidModeliService"
	log := s.log.With(slog.String("op", op))

	NewVidsModeli := partParsers.VidModeliParser(mainStruct, registrator_id)

	// берем из базы имеющие записи и проверяем на дубликаты

	existsVidsModeli, err := s.storage.ListVidModeli(s.ctx)
	if err != nil {
		log.Error("Error selecting vids:", slog.String("err", err.Error()))
		return err
	}
	log.Debug("exist vids: " + strconv.Itoa(len(existsVidsModeli)))

	existsMap := make(map[string]models.VidModeliEntity)
	for _, e := range existsVidsModeli {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.VidModeliEntity

	for _, n := range NewVidsModeli {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateVidModeli(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateVidModeli(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		}
	}

	log.Debug("VidModeli parsing: ", slog.Int("count", len(NewVidsModeli)))
	log.Debug("Duplicated and updated VidModeli: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new VidModeli: ", slog.Int("count", len(toCreate)))

	// for _, val := range NewVidsModeli {
	// 	indexDuplicated := slices.IndexFunc(existsVidsModeli, func(item models.VidModeliEntity) bool {
	// 		return item.Id_1c == val.Id_1c
	// 	})
	// 	if indexDuplicated != -1 {
	// 		log.Debug("Duplicated and updated vids: " + val.Id_1c)
	// 		val.Id = (existsVidsModeli)[indexDuplicated].Id
	// 		if err := s.storage.UpdateVidModeli(s.ctx, val); err != nil {
	// 			log.Error("Error updating vids:", slog.String("err", err.Error()))
	// 			return err
	// 		}
	// 		continue
	// 	}
	// 	if _, err := s.storage.CreateVidModeli(s.ctx, val); err != nil {
	// 		log.Error("Error inserting vids:", slog.String("err", err.Error()))
	// 		return err
	// 	}
	return nil
}
