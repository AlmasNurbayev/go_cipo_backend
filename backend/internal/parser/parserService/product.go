package parserService

import (
	"log/slog"
	"sort"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/partParsers"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func (s *ParserService) ProductService(mainStruct *xmltypes.ImportType,
	registrator_id int64) error {

	op := "ProductService"
	log := s.log.With(slog.String("op", op))

	productGroups, err := s.storage.ListProductGroup(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	productVids, err := s.storage.ListProductVid(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	vidsMideli, err := s.storage.ListVidModeli(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	productDescMappings, err := s.storage.ListProductDescMapping(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	products, err := s.storage.ListProduct(s.ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	NewProducts := partParsers.ProductsParser(mainStruct, registrator_id,
		productGroups, productVids, vidsMideli, productDescMappings)

	// сортируем по первым 20 символам поля id_1c, для хронологии
	sort.Slice(NewProducts, func(i, j int) bool {
		return NewProducts[i].Id_1c[19:] < NewProducts[j].Id_1c[19:]
	})

	existsMap := make(map[string]models.ProductEntity)
	for _, e := range products {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.ProductEntity

	for _, n := range NewProducts {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := s.storage.UpdateProductById1c(s.ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		} else {
			toCreate = append(toCreate, n)
			_, err := s.storage.CreateProduct(s.ctx, n)
			// json, err2 := utils.PrintAsJSON(n)
			// if err2 != nil {
			// 	return err
			// }
			// log.Info(string(*json))
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}

	log.Info("products parsing: ", slog.Int("count", len(NewProducts)))
	log.Info("== Duplicated and updated products: ", slog.Int("count", len(toUpdate)))
	log.Info("== Created new products: ", slog.Int("count", len(toCreate)))

	return nil
}
