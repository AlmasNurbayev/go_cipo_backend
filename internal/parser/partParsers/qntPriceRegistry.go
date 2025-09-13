package partParsers

import (
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
	"github.com/guregu/null/v5"
)

func QntPriceRegistryParser(receiveStruct *xmltypes.OfferType, registrator_id int64,
	operation_date time.Time, stores []models.StoreEntity,
	priceVids []models.PriceVidEntity,
	sizes []models.SizeEntity, products []models.ProductEntity) ([]models.QntPriceRegistryEntity, error) {

	var qntPriceRegistry []models.QntPriceRegistryEntity

	root := receiveStruct.КоммерческаяИнформация.ПакетПредложений.Предложения.Предложение

	for j := range root {

		// готовим мапу продуктов для ускорения поиска
		productMap := make(map[string]models.ProductEntity, len(products))
		for _, product := range products {
			productMap[product.Id_1c] = product
		}
		// ищем в мапе
		subString := utils.GetSubstringIfSymbolExists(root[j].Ид, "#")
		product, found := productMap[subString]
		if !found {
			return nil, errors.New("Not found product in DB " + root[j].Ид)
		}

		// получаем ссылку на размер
		rootSv := root[j].ЗначенияСвойств.ЗначенияСвойства
		var size models.SizeEntity
		for k := range rootSv {
			if !strings.Contains(rootSv[k].Наименование, "Размер") {
				continue
			}
			sizeIndex := slices.IndexFunc(sizes, func(item models.SizeEntity) bool {
				return item.Id_1c == rootSv[k].Значение
			})
			if sizeIndex == -1 {
				return nil, errors.New("not found size in DB " + rootSv[k].Ид)
			}
			size = sizes[sizeIndex]
		}
		if size.Id == 0 {
			return nil, errors.New("not found size in DB")
		}

		// получаем ссылку на склад
		rootStore := root[j].Склад
		var store models.StoreEntity
		var qnt float32
		for k := range rootStore {
			storeIndex := slices.IndexFunc(stores, func(item models.StoreEntity) bool {
				return item.Id_1c == rootStore[k].ИдСклада
			})
			if storeIndex == -1 {
				return nil, errors.New("not found store in DB " + rootStore[k].ИдСклада)
			}
			store = stores[storeIndex]
			qnt = rootStore[k].КоличествоНаСкладе
		}
		if store.Id == 0 {
			return nil, errors.New("not found store in DB")
		}

		// получаем ссылку на вид цены и цену
		rootPriceVid := root[j].Цены.Цена
		var price float32
		var priceZakup float32
		var priceVid models.PriceVidEntity
		for k := range rootPriceVid {
			priceIndex := slices.IndexFunc(priceVids, func(item models.PriceVidEntity) bool {
				// если тип цены совпадает и он активен
				return item.Id_1c == rootPriceVid[k].ИдТипаЦены && item.Active
			})
			if priceIndex != -1 {
				priceVid = priceVids[priceIndex]
				price = rootPriceVid[k].ЦенаЗаЕдиницу
			}
			priceIndexZakup := slices.IndexFunc(priceVids, func(item models.PriceVidEntity) bool {
				// если тип цены совпадает и он активен
				return item.Id_1c == rootPriceVid[k].ИдТипаЦены && item.Is_zakup
			})
			if priceIndexZakup != -1 {
				//priceVidZakup = priceVids[priceIndex]
				priceZakup = utils.RoundFloat32(rootPriceVid[k].ЦенаЗаЕдиницу)
				//fmt.Println("priceZakup=", priceZakup, " for product=", product.Name)
			}
		}
		if priceVid.Id == 0 {
			return nil, errors.New("not found price_vid in DB")
		}
		if price == 0 {
			return nil, errors.New("not found price in DB")
		}

		qntPrice := models.QntPriceRegistryEntity{
			Product_name:        product.Name,
			Product_group_id:    product.Product_group_id,
			Vid_modeli_id:       product.Vid_modeli_id.Int64,
			Nom_vid:             product.Nom_vid,
			Product_create_date: null.TimeFrom(product.Create_date),
			Operation_date:      operation_date,
			Qnt:                 qnt,
			Sum:                 price,
			Sum_zakup:           priceZakup,
			Store_id:            store.Id,
			Price_vid_id:        priceVid.Id,
			Size_id:             size.Id,
			Size_name_1c:        null.StringFrom(size.Name_1c),
			Product_id:          product.Id,
			Registrator_id:      registrator_id,
		}
		//utils.PrintAsJSON(qntPrice)
		qntPriceRegistry = append(qntPriceRegistry, qntPrice)
	}
	return qntPriceRegistry, nil

}
