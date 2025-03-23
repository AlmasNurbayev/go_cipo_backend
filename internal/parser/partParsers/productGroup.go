package partParsers

import (
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func ProductGroupsParser(receiveStruct *xmltypes.ImportType,
	registrator_id int64) []models.ProductsGroupEntity {

	mainStruct := (*receiveStruct)
	var productGroups []models.ProductsGroupEntity

	root := mainStruct.КоммерческаяИнформация.Классификатор.Свойства.Свойство

	for i := 0; i < len(root); i++ {
		if root[i].Наименование == "ТоварнаяГруппа" {

			for j := 0; j < len(root[i].ВариантыЗначений.Справочник); j++ {
				productGroup := models.ProductsGroupEntity{
					Id_1c:          root[i].ВариантыЗначений.Справочник[j].ИдЗначения,
					Name_1c:        root[i].ВариантыЗначений.Справочник[j].Значение,
					Registrator_id: registrator_id,
				}
				productGroups = append(productGroups, productGroup)
			}
		}
	}
	return productGroups
}
