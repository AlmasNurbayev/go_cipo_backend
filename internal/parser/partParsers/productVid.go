package partParsers

import (
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

// ищет в структуре вложенную структуру "Вид обуви" и возвращат ее элементы
func ProductVidsParser(receiveStruct *xmltypes.ImportType, registrator_id int64) []models.ProductVidEntity {

	mainStruct := (*receiveStruct)
	var productVids []models.ProductVidEntity

	root := mainStruct.КоммерческаяИнформация.Классификатор.Группы.Группа

	for i := 0; i < len(root); i++ {
		productVid := models.ProductVidEntity{
			Id_1c:          root[i].Ид,
			Name_1c:        root[i].Наименование,
			Registrator_id: registrator_id,
		}
		productVids = append(productVids, productVid)
	}
	return productVids
}
