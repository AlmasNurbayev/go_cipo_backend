package partParsers

import (
	"slices"
	"strconv"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
	"github.com/guregu/null/v5"
)

func ProductsParser(mainStruct *xmltypes.ImportType, registrator_id int64,
	existsProductGroups []models.ProductsGroupEntity, existsProductVids []models.ProductVidEntity,
	existsVidsModeli []models.VidModeliEntity,
	existsProductDescMappings []models.ProductsDescMappingEntity) []models.ProductEntity {

	var products []models.ProductEntity

	root := mainStruct.КоммерческаяИнформация.Каталог.Товары.Товар

	for i := 0; i < len(root); i++ {

		product_folder := null.StringFrom("")
		var product_vid_id null.Int64
		var description null.String
		var product_desc struct {
			Material_inside   null.String
			Material_podoshva null.String
			Material_up       null.String
			Sex               null.Int16
			Main_color        null.String
			public_web        bool
			Product_group_id  int64
			VidModeli_id      null.Int64
			Kaspi_category    null.String
			Kaspi_in_sale     bool
		}
		var base_ed string
		var nom_vid null.String

		if root[i].Группы.Ид != "" {
			product_folder = null.StringFrom(root[i].Группы.Ид)
		}
		if root[i].БазоваяЕдиница.НаименованиеПолное != "" {
			base_ed = root[i].БазоваяЕдиница.НаименованиеПолное
		}

		if root[i].Описание != "" {
			description = null.StringFrom(root[i].Описание)
		}

		root_rekv := root[i].ЗначенияРеквизитов.ЗначениеРеквизита
		for j := 0; j < len(root_rekv); j++ {
			if root_rekv[j].Наименование == "ВидНоменклатуры" {
				nom_vid = null.StringFrom(root_rekv[j].Значение)
				// product_vid_index := slices.IndexFunc(existsProductVids, func(item models.ProductVidEntity) bool {
				// 	return item.Name_1c == root_rekv[j].Значение
				// })
				// if product_vid_index != -1 {
				// 	product_vid_id = existsProductVids[product_vid_index].Id
				// }
			}
		}

		root_svoistv := root[i].ЗначенияСвойств.ЗначенияСвойства

		// разбираем перечень свойств сопоставляя со справочником Product_desc
		for k := 0; k < len(root_svoistv); k++ {
			for _, val := range existsProductDescMappings {
				if val.Id_1c == root_svoistv[k].Ид {
					if val.Field == "kaspi_category" {
						product_desc.Kaspi_category = null.StringFrom(root_svoistv[k].Значение)
					}
					if val.Field == "kaspi_in_sale" {
						if root_svoistv[k].Значение == "Да" {
							product_desc.Kaspi_in_sale = true
						} else {
							product_desc.Kaspi_in_sale = false
						}
					}
					if val.Field == "material_podoshva" {
						product_desc.Material_podoshva = null.StringFrom(root_svoistv[k].Значение)
					}
					if val.Field == "material_inside" {
						product_desc.Material_inside = null.StringFrom(root_svoistv[k].Значение)
					}
					if val.Field == "material_up" {
						product_desc.Material_up = null.StringFrom(root_svoistv[k].Значение)
					}
					if val.Field == "main_color" {
						product_desc.Main_color = null.StringFrom(root_svoistv[k].Значение)
					}
					if val.Field == "public_web" {
						if root_svoistv[k].Значение == "Да" {
							product_desc.public_web = true
						} else {
							product_desc.public_web = false
						}
					}
					if val.Field == "sex" {
						intSex, err := strconv.Atoi(root_svoistv[k].Значение)
						product_desc.Sex = null.Int16From(int16(intSex))
						if err != nil {
							product_desc.Sex = null.Int16From(0)
						}
					}
					if val.Field == "product_group" {
						// в справочнике Product_desc ищем какое id_1c имеет свойство "ТоварнаяГруппа"
						product_group_index := slices.IndexFunc(existsProductGroups, func(item models.ProductsGroupEntity) bool {
							return item.Id_1c == root_svoistv[k].Значение
						})
						if product_group_index != -1 {
							product_desc.Product_group_id = existsProductGroups[product_group_index].Id
							//product_desc.Product_group_id.Valid = true
						}
					}
					if val.Field == "vidModeli" {
						// в справочнике Product_desc ищем какое id_1c имеет свойство "Виды"
						vid_index := slices.IndexFunc(existsVidsModeli, func(item models.VidModeliEntity) bool {
							return item.Id_1c == root_svoistv[k].Значение
						})
						if vid_index != -1 {
							product_desc.VidModeli_id = null.IntFrom(existsVidsModeli[vid_index].Id)
						}
					}
				}
			}
		}

		newProduct := models.ProductEntity{
			Id_1c:             root[i].Ид,
			Name_1c:           root[i].Наименование,
			Name:              root[i].Наименование,
			Registrator_id:    registrator_id,
			Artikul:           root[i].Артикул,
			Description:       description,
			Base_ed:           base_ed,
			Product_group_id:  product_desc.Product_group_id,
			Product_folder:    product_folder,
			Product_vid_id:    product_vid_id,
			Material_podoshva: product_desc.Material_podoshva,
			Material_inside:   product_desc.Material_inside,
			Material_up:       product_desc.Material_up,
			Sex:               product_desc.Sex,
			Main_color:        product_desc.Main_color,
			Public_web:        product_desc.public_web,
			Vid_modeli_id:     product_desc.VidModeli_id,
			Nom_vid:           nom_vid,
			Kaspi_category:    product_desc.Kaspi_category,
			Kaspi_in_sale:     product_desc.Kaspi_in_sale,
		}
		products = append(products, newProduct)
	}
	return products

}
