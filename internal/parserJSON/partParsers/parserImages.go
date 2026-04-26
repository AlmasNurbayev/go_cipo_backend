package partParsers

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/gofiber/fiber/v3/log"
	"github.com/guregu/null/v5"
	"github.com/kr/pretty"
)

type StorageI interface {
	ListImageRegistry(context.Context) ([]models.ImageRegistryEntity, error)
	UpdateImageRegistryByName(context.Context, models.ImageRegistryEntity) error
	CreateImageRegistry(context.Context, models.ImageRegistryEntity) (int64, error)
	ListProduct(context.Context) ([]models.ProductEntity, error)
}

func ImageRegistryParser(ctx context.Context, storage StorageI, Log *slog.Logger, data dto.StockJSON, registrator_id int64,
	newPath string) error {

	op := "parserJSON.ImageRegistryParser"
	Log = Log.With(slog.String("op", op))

	existsProducts, err := storage.ListProduct(ctx)
	if err != nil {
		return err
	}

	time := time.Now()
	var newImages []models.ImageRegistryEntity

	for _, itemQnt := range data.Qnt {
		if len(itemQnt.Images) == 0 {
			continue
		}
		is_exists_product := slices.IndexFunc(existsProducts, func(item models.ProductEntity) bool {
			return item.Id_1c == itemQnt.NomGUID
		})
		if is_exists_product == -1 {
			return errors.New("in JSON found product but not exists in DB " + itemQnt.NomName)
		}
		countImagesInProduct := 0

		root_images := itemQnt.Images
		for _, imageItem := range root_images {
			fileInfo, err := os.Stat(newPath + "/" + imageItem)
			if err != nil {
				return errors.New("Error getting file information: " + err.Error())
			}
			full_name := newPath + "/" + imageItem
			var is_main bool
			if countImagesInProduct == 0 {
				is_main = true
			} else {
				is_main = false
			}
			part_name := "product_images/" + imageItem

			image := models.ImageRegistryEntity{
				Full_name:          part_name,
				Name:               filepath.Base(full_name),
				Path:               strings.TrimPrefix(filepath.Dir(full_name), "../assets/"),
				Size:               int(fileInfo.Size()),
				Resolution:         null.StringFromPtr(nil),
				Active_change_date: time,
				Active:             true,
				Main_change_date:   time,
				Main:               is_main,
				Registrator_id:     registrator_id,
				Product_id:         existsProducts[is_exists_product].Id,

				Operation_date: time,
			}
			newImages = append(newImages, image)
			countImagesInProduct = countImagesInProduct + 1
		}

	}

	// считываем все уже имющиеся записи в products
	existsImages, err := storage.ListImageRegistry(ctx)
	if err != nil {
		log.Error("Error selecting images: ", slog.String("err", err.Error()))
		return err
	}

	Log.Debug("image exists count: ", slog.Int("count", len(existsImages)))

	existsMap := make(map[string]models.ImageRegistryEntity)
	for _, e := range existsImages {
		existsMap[e.Name] = e
	}

	var toCreate, toUpdate []models.ImageRegistryEntity

	for _, n := range newImages {
		if _, exists := existsMap[n.Name]; exists {
			// Элемент есть → Обновляем (если изменился)
			toUpdate = append(toUpdate, n)
			err := storage.UpdateImageRegistryByName(ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		} else {
			toCreate = append(toCreate, n)
			_, err := storage.CreateImageRegistry(ctx, n)
			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}

	Log.Debug("images parsing: ", slog.Int("count", len(newImages)))
	Log.Debug("Duplicated and updated images: ", slog.Int("count", len(toUpdate)))
	Log.Info("Created new images: ", slog.Int("count", len(toCreate)))

	pretty.Log(newImages[0])

	return nil
}
