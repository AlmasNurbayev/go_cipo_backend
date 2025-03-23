package partParsers

import (
	"time"

	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
)

func PriceVidParser(receiveStruct *xmltypes.OfferType, registrator_id int64) []models.PriceVidEntity {

	var priceVids []models.PriceVidEntity
	time := time.Now()

	root := receiveStruct.КоммерческаяИнформация.ПакетПредложений.ТипыЦен.ТипЦены

	for j := range root {
		//var vid T

		priceVid := models.PriceVidEntity{
			Id_1c:              root[j].Ид,
			Name_1c:            root[j].Наименование,
			Active_change_date: time,
			Registrator_id:     registrator_id,
		}
		if strings.Contains(strings.ToLower(root[j].Наименование), "розничная") {
			priceVid.Active = true
		} else {
			priceVid.Active = false
		}

		priceVids = append(priceVids, priceVid)
	}
	return priceVids
}
