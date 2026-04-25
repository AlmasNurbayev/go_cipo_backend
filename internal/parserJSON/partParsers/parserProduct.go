package partParsers

import (
	"context"
	"errors"
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/guregu/null/v5"
	"github.com/kr/pretty"
)

type StorageP interface {
	ListProductGroup(context.Context) ([]models.ProductsGroupEntity, error)
	ListProductVid(context.Context) ([]models.ProductVidEntity, error)
	ListVidModeli(context.Context) ([]models.VidModeliEntity, error)
	ListProductDescMapping(context.Context) ([]models.ProductsDescMappingEntity, error)
	ListProduct(context.Context) ([]models.ProductEntity, error)
	ListBrend(context.Context) ([]models.BrendEntity, error)
	CreateProduct(context.Context, models.ProductEntity) (int64, error)
	UpdateProductById1c(context.Context, models.ProductEntity) error
}

func ParserProduct(Log *slog.Logger, ctx context.Context, storage StorageP, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.partParsers.ParserProduct"
	log := Log.With(slog.String("op", op))

	productGroups, err := storage.ListProductGroup(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	productVids, err := storage.ListProductVid(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	vidsMideli, err := storage.ListVidModeli(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	productDescMappings, err := storage.ListProductDescMapping(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	brends, err := storage.ListBrend(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var products []models.ProductEntity

	root := data.Qnt

	for _, item := range root {

		//product_folder := null.StringFrom("")
		var product_vid_id, brend_id null.Int64
		var description null.String
		var product_group_id int64
		var product_desc struct {
			Material_inside   null.String
			Material_podoshva null.String
			Material_up       null.String
			Sex               null.Int16
			Main_color        null.String
			public_web        bool

			VidModeli_id   null.Int64
			Kaspi_category null.String
			Kaspi_in_sale  bool
		}
		var base_ed string
		var nom_vid null.String

		// if root[i].Группы.Ид != "" {
		// 	product_folder = null.StringFrom(root[i].Группы.Ид)
		// }
		base_ed = item.EdIzmName
		description = null.StringFrom(item.NomDescription)

		for _, productVid := range productVids {
			if productVid.Id_1c == item.VidGUID {
				product_vid_id = null.IntFrom(productVid.Id)
				nom_vid = null.StringFrom(productVid.Name_1c)
			}
		}

		for _, productGroup := range productGroups {
			if productGroup.Id_1c == item.TovarGroupGUID {
				product_group_id = productGroup.Id
			}
		}

		for _, brend := range brends {
			if brend.Id_1c == item.MarkaGUID {
				brend_id = null.IntFrom(brend.Id)
			}
		}

		//for _, root_prop := range item.AdditionalProperties {
		//if strings.Contains(root_prop.NameProperty, "ВидНоменклатуры") {
		//nom_vid = null.StringFrom(root_prop.StrValueProperty)
		// product_vid_index := slices.IndexFunc(existsProductVids, func(item models.ProductVidEntity) bool {
		// 	return item.Name_1c == root_rekv[j].Значение
		// })
		// if product_vid_index != -1 {
		// 	product_vid_id = existsProductVids[product_vid_index].Id
		// }
		//}
		//}

		// 	root_svoistv := root[i].ЗначенияСвойств.ЗначенияСвойства

		// 	// разбираем перечень свойств сопоставляя со справочником Product_desc
		// 	for k := 0; k < len(root_svoistv); k++ {
		// 		for _, val := range existsProductDescMappings {
		// 			if val.Id_1c == root_svoistv[k].Ид {
		// 				if val.Field == "kaspi_category" {
		// 					product_desc.Kaspi_category = null.StringFrom(root_svoistv[k].Значение)
		// 				}
		// 				if val.Field == "kaspi_in_sale" {
		// 					if root_svoistv[k].Значение == "Да" {
		// 						product_desc.Kaspi_in_sale = true
		// 					} else {
		// 						product_desc.Kaspi_in_sale = false
		// 					}
		// 				}
		// 				if val.Field == "material_podoshva" {
		// 					product_desc.Material_podoshva = null.StringFrom(root_svoistv[k].Значение)
		// 				}
		// 				if val.Field == "material_inside" {
		// 					product_desc.Material_inside = null.StringFrom(root_svoistv[k].Значение)
		// 				}
		// 				if val.Field == "material_up" {
		// 					product_desc.Material_up = null.StringFrom(root_svoistv[k].Значение)
		// 				}
		// 				if val.Field == "main_color" {
		// 					product_desc.Main_color = null.StringFrom(root_svoistv[k].Значение)
		// 				}
		// 				if val.Field == "public_web" {
		// 					if root_svoistv[k].Значение == "Да" {
		// 						product_desc.public_web = true
		// 					} else {
		// 						product_desc.public_web = false
		// 					}
		// 				}
		// 				if val.Field == "sex" {
		// 					intSex, err := strconv.Atoi(root_svoistv[k].Значение)
		// 					product_desc.Sex = null.Int16From(int16(intSex))
		// 					if err != nil {
		// 						product_desc.Sex = null.Int16From(0)
		// 					}
		// 				}
		// 				if val.Field == "product_group" {
		// 					// в справочнике Product_desc ищем какое id_1c имеет свойство "ТоварнаяГруппа"
		// 					product_group_index := slices.IndexFunc(existsProductGroups, func(item models.ProductsGroupEntity) bool {
		// 						return item.Id_1c == root_svoistv[k].Значение
		// 					})
		// 					if product_group_index != -1 {
		// 						product_desc.Product_group_id = existsProductGroups[product_group_index].Id
		// 						//product_desc.Product_group_id.Valid = true
		// 					}
		// 				}
		// 				if val.Field == "vidModeli" {
		// 					// в справочнике Product_desc ищем какое id_1c имеет свойство "Виды"
		// 					vid_index := slices.IndexFunc(existsVidsModeli, func(item models.VidModeliEntity) bool {
		// 						return item.Id_1c == root_svoistv[k].Значение
		// 					})
		// 					if vid_index != -1 {
		// 						product_desc.VidModeli_id = null.IntFrom(existsVidsModeli[vid_index].Id)
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}

		if product_group_id == 0 {
			log.Debug(item.NomArticle)
			log.Warn("product group not found: ", slog.String("productName", item.NomName))
			continue
			//return errors.New("product group not found")
		}

		if !product_vid_id.Valid {
			log.Error("product vid not found: ", slog.String("productName", item.NomName))
			return errors.New("product vid not found")
		}

		newProduct := models.ProductEntity{
			Id_1c:            item.NomGUID,
			Name_1c:          item.NomName,
			Name:             item.NomName,
			Registrator_id:   registrator_id,
			Artikul:          item.NomArticle,
			Description:      description,
			Base_ed:          base_ed,
			Brend_id:         brend_id,
			Product_group_id: product_group_id,
			//Product_folder:    product_folder,
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
	// save products

	pretty.Println("productGroups:", productGroups)
	pretty.Println("productVids:", productVids)
	pretty.Println("vidsMideli:", vidsMideli)
	pretty.Println("productDescMappings:", productDescMappings)
	pretty.Println("products:", products[10])
	return nil
}
