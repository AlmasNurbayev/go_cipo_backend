package partParsers

import (
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func SizeParser(receiveStruct *xmltypes.OfferType, registrator_id int64) []models.SizeEntity {
	mainStruct := (*receiveStruct)
	var sizes []models.SizeEntity

	root := mainStruct.КоммерческаяИнформация.Классификатор.Свойства.Свойство

	// TODO - в этой ветке стуктуры единственный массив, может быть не проверять Ид
	if root.Ид == "a001d8e3-a3b3-11ed-b0d2-50ebf624c538" {

		for j := range root.ВариантыЗначений.Справочник {
			//var vid T
			size := models.SizeEntity{
				Id_1c:          root.ВариантыЗначений.Справочник[j].ИдЗначения,
				Name_1c:        root.ВариантыЗначений.Справочник[j].Значение,
				Registrator_id: registrator_id,
			}
			sizes = append(sizes, size)
		}
	}
	return sizes
}
