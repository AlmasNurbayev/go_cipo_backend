package partParsers

import (
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

// ищет в структуре вложенную структуру "Вид товаров" и возвращат ее элементы
func VidModeliParser(receiveStruct *xmltypes.ImportType, registrator_id int64) []models.VidModeliEntity {

	mainStruct := (*receiveStruct)
	var vids []models.VidModeliEntity

	root := mainStruct.КоммерческаяИнформация.Классификатор.Свойства.Свойство

	for i := 0; i < len(root); i++ {
		if root[i].Наименование == "ВидМодели" {

			for j := 0; j < len(root[i].ВариантыЗначений.Справочник); j++ {
				//var vid T
				vid := models.VidModeliEntity{
					Id_1c:          root[i].ВариантыЗначений.Справочник[j].ИдЗначения,
					Name_1c:        root[i].ВариантыЗначений.Справочник[j].Значение,
					Registrator_id: registrator_id,
				}
				vids = append(vids, vid)
			}
		}
	}
	return vids
}
