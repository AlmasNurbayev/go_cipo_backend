package parserService

import (
	"log/slog"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) ProductGroupService(mainStruct *xmltypes.ImportType,
	registrator_id int64) error {

	op := "ProductGroupService"
	log := s.log.With(slog.String("op", op))

	NewProductGroups := partParsers.ProductGroupsParser(mainStruct, registrator_id)

	// берем из базы имеющие записи и проверяем на дубликаты
	existsProductGroups, err := s.storage.ListProductGroup(s.ctx)
	if err != nil {
		log.Error("Error selecting product_groups:", slog.String("err", err.Error()))
		return err
	}
	log.Debug("exist product_groups: " + strconv.Itoa(len(existsProductGroups)))

	existsMap := make(map[string]models.ProductsGroupEntity)
	for _, e := range existsProductGroups {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.ProductsGroupEntity

	for _, n := range NewProductGroups {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateProductGroup(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		} else {
			// Элемента нет → Добавляем
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateProductGroup(s.ctx, n)
			if err != nil {
				s.log.Error(err.Error())
			}
		}
	}

	log.Debug("ProductGroup parsing: ", slog.Int("count", len(NewProductGroups)))
	log.Debug("Duplicated and updated ProductGroup: ", slog.Int("count", len(toUpdate)))
	log.Info("Created new ProductGroup: ", slog.Int("count", len(toCreate)))

	return nil
}
