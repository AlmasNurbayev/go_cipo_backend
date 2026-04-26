package partParsers

import (
	"context"
	"errors"
	"log/slog"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/guregu/null/v5"
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
	vidsModeli, err := storage.ListVidModeli(ctx)
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

	currentProducts, err := storage.ListProduct(ctx)
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

		for _, mappingItem := range productDescMappings {
			for _, root_prop := range item.AdditionalProperties {
				// здесь ищем не по ID свойства в productDescMapping, а просто по наименованию свойства
				if strings.Contains(root_prop.NameProperty, "ВидМодели") {
					idx := slices.IndexFunc(vidsModeli, func(e models.VidModeliEntity) bool {
						return e.Id_1c == root_prop.GUIDValue
					})
					if idx != -1 {
						product_desc.VidModeli_id = null.IntFrom(vidsModeli[idx].Id)
					}
				}
				// далее обязательно нужно соответствие ГУИД из productDescMapping
				if mappingItem.Id_1c != root_prop.GUIDProperty {
					continue
				}
				if mappingItem.Field == "material_podoshva" {
					product_desc.Material_podoshva = null.StringFrom(root_prop.StrValueProperty)
				}
				if mappingItem.Field == "material_inside" {
					product_desc.Material_inside = null.StringFrom(root_prop.StrValueProperty)
				}
				if mappingItem.Field == "material_up" {
					product_desc.Material_up = null.StringFrom(root_prop.StrValueProperty)
				}
				if mappingItem.Field == "main_color" {
					product_desc.Main_color = null.StringFrom(root_prop.StrValueProperty)
				}
				if mappingItem.Field == "kaspi_category" {
					product_desc.Kaspi_category = null.StringFrom(root_prop.StrValueProperty)
				}
				if mappingItem.Field == "kaspi_in_sale" {
					if root_prop.StrValueProperty == "Да" {
						product_desc.Kaspi_in_sale = true
					} else {
						product_desc.Kaspi_in_sale = false
					}
				}
				if mappingItem.Field == "sex" {
					intSex, err := strconv.Atoi(root_prop.StrValueProperty)
					if err == nil {
						product_desc.Sex = null.Int16From(int16(intSex))
					}
				}
			}
		}

		for _, item_rekv := range item.AdditionalRekvizits {
			if item_rekv.StrValueRekv == "" {
				continue
			}
			if strings.Contains(item_rekv.NameRekv, "ВыгружатьВеб") {
				if item_rekv.StrValueRekv == "Да" {
					product_desc.public_web = true
				} else {
					product_desc.public_web = false
				}
			}
		}

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

	//pretty.Println("products:", products[0])
	//return nil

	// сортируем по части символов поля id_1c, для хронологии
	sort.Slice(products, func(i, j int) bool {
		//log.Debug(NewProducts[i].Id_1c[14:23])
		return products[i].Id_1c[14:23] < products[j].Id_1c[14:23]
	})

	existsMap := make(map[string]models.ProductEntity)
	for _, e := range currentProducts {
		existsMap[e.Id_1c] = e
	}

	var toCreate, toUpdate []models.ProductEntity

	for _, n := range products {
		if _, exists := existsMap[n.Id_1c]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateProductById1c(ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		} else {
			toCreate = append(toCreate, n)
			_, err := storage.CreateProduct(ctx, n)
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

	log.Info("products parsing: ", slog.Int("count", len(products)))
	log.Info("== Duplicated and updated products: ", slog.Int("count", len(toUpdate)))
	log.Info("== Created new products: ", slog.Int("count", len(toCreate)))

	// pretty.Println("productGroups:", productGroups)
	// pretty.Println("productVids:", productVids)
	// pretty.Println("vidsMideli:", vidsMideli)
	// pretty.Println("productDescMappings:", productDescMappings)
	// pretty.Println("products:", products[0])
	return nil
}
