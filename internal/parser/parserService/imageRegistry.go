package parserService

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) ImageRegistryService(mainStruct *xmltypes.ImportType, registrator_id int64, newPath string) error {
	op := "parserService.ImageRegistryService"
	log := s.log.With(slog.String("op", op))

	existsProducts, err := s.storage.ListProduct(s.ctx)
	if err != nil {
		log.Error("Error selecting products: ", slog.String("err", err.Error()))
		return err
	}

	newImages, err := partParsers.ImageRegistryParser(*mainStruct, registrator_id, existsProducts, newPath)
	if err != nil {
		log.Error("Error parsing images: ", slog.String("err", err.Error()))
		return err
	}

	// считываем все уже имющиеся записи в products
	existsImages, err := s.storage.ListImageRegistry(s.ctx)
	if err != nil {
		log.Error("Error selecting images: ", slog.String("err", err.Error()))
		return err
	}

	log.Debug("image exists count: ", slog.Int("count", len(existsImages)))

	existsMap := make(map[string]models.ImageRegistryEntity)
	for _, e := range existsImages {
		existsMap[e.Name] = e
	}

	var toCreate, toUpdate []models.ImageRegistryEntity

	for _, n := range newImages {
		if _, exists := existsMap[n.Name]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateImageRegistryByName(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		} else {
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateImageRegistry(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}

	log.Debug("images parsing: ", slog.Int("count", len(newImages)))
	log.Debug("Duplicated and updated images: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new images: ", slog.Int("count", len(toCreate)))

	return nil
}
