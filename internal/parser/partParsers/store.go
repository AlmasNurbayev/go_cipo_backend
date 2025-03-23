package partParsers

import (
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func StoreParser(receiveStruct *xmltypes.OfferType, registrator_id int64) []models.StoreEntity {
	mainStruct := (*receiveStruct)
	var stores []models.StoreEntity

	root := mainStruct.КоммерческаяИнформация.ПакетПредложений.Склады.Склад

	for j := range root {
		//var vid T
		store := models.StoreEntity{
			Id_1c:          root[j].Ид,
			Name_1c:        root[j].Наименование,
			Registrator_id: registrator_id,
		}
		stores = append(stores, store)
	}

	return stores
}
