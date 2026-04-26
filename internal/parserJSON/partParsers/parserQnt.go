package partParsers

import (
	"context"
	"errors"
	"log/slog"
	"slices"
	"strings"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/guregu/null/v5"
)

type storageQnt interface {
	ListPriceVid(ctx context.Context) ([]models.PriceVidEntity, error)
	ListSize(ctx context.Context) ([]models.SizeEntity, error)
	ListStore(ctx context.Context) ([]models.StoreEntity, error)
	ListProduct(context.Context) ([]models.ProductEntity, error)
	ListProductVid(context.Context) ([]models.ProductVidEntity, error)
	CreateQntPriceRegistry(context.Context, models.QntPriceRegistryEntity) (int64, error)
}

func ParserQnt(Log *slog.Logger, ctx context.Context, storage storageQnt, data dto.StockJSON, registrator_id int64) error {
	op := "parserJSON.parserQnt"
	Log = Log.With(slog.String("op", op))

	location := time.Local
	layout := "2006-01-02T15:04:05"
	operation_date, err := time.ParseInLocation(layout, data.ExportDate, location)
	if err != nil {
		Log.Error("Error parsing time:", slog.String("err", err.Error()))
		return err
	}

	priceVids, err := storage.ListPriceVid(ctx)
	if err != nil {
		Log.Error("error list price vids: ", slog.String("error", err.Error()))
		return err
	}
	indexPrice := slices.IndexFunc(priceVids, func(priceVid models.PriceVidEntity) bool {
		return strings.Contains(priceVid.Name_1c, "Розничная")
	})
	if indexPrice == -1 {
		Log.Error("error price Розничная not found: ", slog.String("error", err.Error()))
		return err
	}
	roznPrice := priceVids[indexPrice].Id

	sizes, err := storage.ListSize(ctx)
	if err != nil {
		Log.Error("error list sizes: ", slog.String("error", err.Error()))
		return err
	}

	stores, err := storage.ListStore(ctx)
	if err != nil {
		Log.Error("error list stores: ", slog.String("error", err.Error()))
		return err
	}

	productVids, err := storage.ListProductVid(ctx)
	if err != nil {
		Log.Error("error list product vids: ", slog.String("error", err.Error()))
		return err
	}

	products, err := storage.ListProduct(ctx)
	if err != nil {
		Log.Error("error list products: ", slog.String("error", err.Error()))
		return err
	}

	root := data.Qnt
	newQnts := []models.QntPriceRegistryEntity{}

	for _, qnt := range root {
		// явно пропускаем
		if qnt.NomName == "Тестовый" {
			Log.Debug("skipped product: Тестовый")
			continue
		}

		for _, stockItem := range qnt.Stocks {
			if stockItem.WarehouseGUID == "" || stockItem.StockGUID == "" ||
				stockItem.CharGUID == "" {
				return errors.New("error: " + qnt.NomName + " contains in Qnt not correct values")

			}
			if stockItem.Quantity == 0 || stockItem.Price == 0 {
				continue
			}
			productArr_id := slices.IndexFunc(products, func(t models.ProductEntity) bool {
				return t.Id_1c == qnt.NomGUID
			})
			if productArr_id == -1 {
				return errors.New("guid product not found " + qnt.NomGUID)
			}
			product_id := products[productArr_id].Id

			if !products[productArr_id].Vid_modeli_id.Valid {
				return errors.New("vid modeli NULL " + products[productArr_id].Name)
			}
			if products[productArr_id].Vid_modeli_id.Int64 == 0 {
				return errors.New("vid modeli is 0 " + products[productArr_id].Name)
			}
			product_vid_id := int64(0)
			if products[productArr_id].Product_vid_id.Valid {
				product_vid_id = int64(slices.IndexFunc(productVids, func(t models.ProductVidEntity) bool {
					return t.Id == products[productArr_id].Product_vid_id.Int64
				}))
			}
			product_vid := productVids[product_vid_id].Name_1c

			product_group_id := products[productArr_id].Product_group_id

			storeArr_id := slices.IndexFunc(stores, func(t models.StoreEntity) bool {
				return t.Id_1c == stockItem.StockGUID
			})
			if storeArr_id == -1 {
				return errors.New("guid store not found")
			}
			store_id := stores[storeArr_id].Id

			sizeArr_id := slices.IndexFunc(sizes, func(t models.SizeEntity) bool {
				return t.Id_1c == stockItem.CharGUID
			})
			if sizeArr_id == -1 {
				return errors.New("guid size not found")
			}
			size_id := sizes[sizeArr_id].Id

			newQnt := models.QntPriceRegistryEntity{
				Registrator_id:      registrator_id,
				Price_vid_id:        roznPrice,
				Size_id:             size_id,
				Store_id:            store_id,
				Product_id:          product_id,
				Product_create_date: null.NewTime(products[productArr_id].Create_date, true),
				Product_name:        products[productArr_id].Name,
				Vid_modeli_id:       products[productArr_id].Vid_modeli_id.Int64,
				Product_group_id:    product_group_id,
				Operation_date:      operation_date,
				Nom_vid:             null.StringFrom(product_vid),
				Size_name_1c:        null.StringFrom(sizes[sizeArr_id].Name_1c),
				Qnt:                 float32(stockItem.Quantity),
				Sum:                 float32(stockItem.Price),
				Sum_zakup:           float32(0),
				Barcode:             null.StringFrom(stockItem.Barcode),
			}

			_, err = storage.CreateQntPriceRegistry(ctx, newQnt)
			if err != nil {
				return err
			}

			newQnts = append(newQnts, newQnt)
		}
	}

	Log.Info("created new qnt_price records: ", slog.Int("count", len(newQnts)))

	return nil
}
