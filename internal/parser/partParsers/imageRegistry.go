package partParsers

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/parser/xmltypes"
	"github.com/guregu/null/v5"
)

func ImageRegistryParser(mainStruct xmltypes.ImportType, registrator_id int64,
	existsProducts []models.ProductEntity, newPath string) ([]models.ImageRegistryEntity,
	error) {

	time := time.Now()
	var images []models.ImageRegistryEntity

	root := mainStruct.КоммерческаяИнформация.Каталог.Товары.Товар

	for productIndex := 0; productIndex < len(root); productIndex++ {
		if root[productIndex].Ид == "" {
			continue
		}
		is_exists_product := slices.IndexFunc(existsProducts, func(item models.ProductEntity) bool {
			return item.Id_1c == root[productIndex].Ид
		})
		if is_exists_product == -1 {
			return nil, errors.New("in XML found product but not exists in DB " + root[productIndex].Ид)
		}
		countImagesInProduct := 0

		root_images := root[productIndex].Картинка
		for imageIndex := 0; imageIndex < len(root_images); imageIndex++ {
			full_name := strings.ReplaceAll(root_images[imageIndex], "import_files", "product_images")
			fileInfo, err := os.Stat("assets/" + full_name)
			if err != nil {
				return nil, errors.New("Error getting file information: " + err.Error())
			}
			var is_main bool
			if countImagesInProduct == 0 {
				is_main = true
			} else {
				is_main = false
			}
			image := models.ImageRegistryEntity{
				Full_name:          full_name,
				Name:               filepath.Base(full_name),
				Path:               filepath.Dir(full_name),
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
			images = append(images, image)
			countImagesInProduct = countImagesInProduct + 1
		}
	}
	return images, nil
}
