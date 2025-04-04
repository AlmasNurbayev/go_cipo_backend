package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"slices"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/clients"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/dto"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/errorsShare"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/lib/utils"
	"github.com/AlmasNurbayev/go_cipo_backend/internal/models"
	"github.com/jinzhu/copier"
)

func (s *Service) KaspiAddCategory(ctx context.Context, data dto.KaspiAddCategoryRequest) (dto.KaspiAddCategoryResponse, error) {

	newCategory := models.KaspiCategoriesEntity{
		Name_kaspi: data.Name_kaspi,
		// Title_kaspi:  data.Title_kaspi, нужно получить из списка категорий
	}

	response := dto.KaspiAddCategoryResponse{}
	response.Name_kaspi = data.Name_kaspi

	organizations, err := s.postgresStorage.ListKaspiOrganization(ctx)
	if err != nil {
		return response, err
	}

	var token string
	for _, organization := range organizations {
		if organization.Id == data.OrganizationId {
			token, err = utils.DecryptToken(s.cfg.Auth.SECRET_BYTE, organization.Kaspi_api_token_hash)
			if err != nil {
				return response, errorsShare.ErrInternalError.Error
			}
			fmt.Println("organization", organization)
		}
	}
	if token == "" {
		// если что-то не так с токеном
		return response, errorsShare.ErrInternalError.Error
	}

	// получаем список категорий и проверяем
	statusCategories, bodyCategories, _ := clients.KaspiGetCategories(s.cfg, s.log, token)
	type category struct {
		code  string
		title string
	}
	categories := []category{}

	if statusCategories == 200 {
		var categoriesTemp []map[string]any
		// Парсим JSON
		err := json.Unmarshal([]byte(bodyCategories), &categoriesTemp)
		if err != nil {
			fmt.Println("err:", err)
			return response, err
		}
		for _, item := range categoriesTemp {
			categories = append(categories, category{code: item["code"].(string), title: item["title"].(string)})
		}
	} else {
		fmt.Println("===5")
		//s.log.Error("Error", slog.String("err", errCategories.Error()))
		return response, fmt.Errorf("categories response %d", statusCategories)
	}
	titleIndex := slices.IndexFunc(categories, func(category category) bool {
		return category.code == data.Name_kaspi
	})
	if titleIndex == -1 {
		return response, errors.New("category code not found")
	} else {
		newCategory.Title_kaspi = categories[titleIndex].title
	}

	//получаем список обязательных атрибутов
	status, body, err := clients.KaspiGetAttributes(s.cfg, s.log, data.Name_kaspi, token)
	if err != nil {
		return response, err
	}
	var attributes []map[string]any
	//var mandatoryAttributes []map[string]any
	if status == 200 {
		// Парсим JSON
		err := json.Unmarshal([]byte(body), &attributes)
		if err != nil {
			fmt.Println("err:", err)
			return response, err
		}
		if len(attributes) > 0 {
			newCategory.Attributes_list = attributes
		}
		// for _, item := range attributes {
		// 	if item["mandatory"].(bool) {
		// 		mandatoryAttributes = append(mandatoryAttributes, item)
		// 	}
		// }
	} else {
		return response, errors.New("get list attributes of category response not 200")
	}
	str, _ := utils.PrintAsJSON(attributes)

	fmt.Println("attributes", string(*str))

	values := clients.KaspiGetAllValues(s.cfg, s.log, data.Name_kaspi, token)
	// strValues, _ := utils.PrintAsJSON(values)
	// fmt.Println("strValues", string(*strValues))
	for _, attribute := range values {
		slices.Sort(attribute.Items)
		switch attribute.ValueName {
		case "Shoes*Size":
			newCategory.Size_kaspi = attribute.Items
			if len(newCategory.Size_kaspi) != 0 {
				newCategory.First_size = newCategory.Size_kaspi[0]
				newCategory.Last_size = newCategory.Size_kaspi[len(newCategory.Size_kaspi)-1]
			}
		case "Shoes*Season":
			newCategory.Season_kaspi = attribute.Items
		case "Shoes*Material":
			newCategory.Material_kaspi = attribute.Items
		case "Shoes*Gender":
			newCategory.Gender_kaspi = attribute.Items
		case "Shoes*Colour":
			newCategory.Colour_kaspi = attribute.Items
		case "Shoes*Model":
			newCategory.Model_kaspi = attribute.Items
		}
	}

	// // получаем список размеров
	// statusSize, bodySize, err := clients.KaspiGetValues(s.cfg, s.log, "Shoes*Size", data.Name_kaspi, token)
	// if err != nil {
	// 	return response, err
	// }
	// sizes := []string{}
	// if statusSize == 200 {
	// 	var rawData []map[string]interface{}
	// 	if err := json.Unmarshal([]byte(bodySize), &rawData); err != nil {
	// 		return response, err
	// 	}
	// 	for _, obj := range rawData {
	// 		if size, ok := obj["name"].(string); ok {
	// 			sizes = append(sizes, size)
	// 		}
	// 	}
	// 	sort.Strings(sizes)
	// 	newCategory.Size_kaspi = sizes
	// }

	// // получаем список сезонов
	// statusSeason, bodySeason, err := clients.KaspiGetValues(s.cfg, s.log, "Shoes*Season", data.Name_kaspi, token)
	// if err != nil {
	// 	return response, err
	// }
	// seasons := []string{}
	// if statusSeason == 200 {
	// 	var rawData []map[string]interface{}
	// 	if err := json.Unmarshal([]byte(bodySeason), &rawData); err != nil {
	// 		return response, err
	// 	}
	// 	for _, obj := range rawData {
	// 		if season, ok := obj["name"].(string); ok {
	// 			seasons = append(seasons, season)
	// 		}
	// 	}
	// 	sort.Strings(seasons)
	// 	newCategory.Season_kaspi = seasons
	// }

	// // получаем список цветов
	// statusColors, bodyColors, err := clients.KaspiGetValues(s.cfg, s.log, "Shoes*Colour", data.Name_kaspi, token)
	// if err != nil {
	// 	return response, err
	// }
	// colours := []string{}
	// if statusColors == 200 {
	// 	var rawData []map[string]interface{}
	// 	if err := json.Unmarshal([]byte(bodyColors), &rawData); err != nil {
	// 		return response, err
	// 	}
	// 	for _, obj := range rawData {
	// 		if color, ok := obj["name"].(string); ok {
	// 			colours = append(colours, color)
	// 		}
	// 	}
	// 	sort.Strings(colours)
	// 	newCategory.Colour_kaspi = colours
	// }

	// // получаем список цветов
	// statusModels, bodyModels, err := clients.KaspiGetValues(s.cfg, s.log, "Shoes*Model", data.Name_kaspi, token)
	// if err != nil {
	// 	return response, err
	// }
	// models := []string{}
	// if statusModels == 200 {
	// 	var rawData []map[string]interface{}
	// 	if err := json.Unmarshal([]byte(bodyModels), &rawData); err != nil {
	// 		return response, err
	// 	}
	// 	for _, obj := range rawData {
	// 		if model, ok := obj["name"].(string); ok {
	// 			models = append(models, model)
	// 		}
	// 	}
	// 	sort.Strings(colours)
	// 	newCategory.Model_kaspi = models
	// }

	// // получаем список полов
	// statusGenders, bodyGenders, err := clients.KaspiGetValues(s.cfg, s.log, "Shoes*Gender", data.Name_kaspi, token)
	// if err != nil {
	// 	return response, err
	// }
	// models := []string{}
	// if statusModels == 200 {
	// 	var rawData []map[string]interface{}
	// 	if err := json.Unmarshal([]byte(bodyModels), &rawData); err != nil {
	// 		return response, err
	// 	}
	// 	for _, obj := range rawData {
	// 		if model, ok := obj["name"].(string); ok {
	// 			models = append(models, model)
	// 		}
	// 	}
	// 	sort.Strings(colours)
	// 	newCategory.Model_kaspi = models
	// }

	result, err := s.postgresStorage.CreateKaspiCategory(ctx, newCategory)
	if err != nil {
		s.log.Error("", slog.String("err", err.Error()))
		return response, err
	}

	err = copier.Copy(&response, &result)
	if err != nil {
		s.log.Error("", slog.String("err", err.Error()))
		return response, err
	}

	return response, nil
}

func (s *Service) KaspiListCategory(ctx context.Context) (dto.KaspiListCategoryResponse, error) {
	op := "services.KaspiListCategory"
	log := s.log.With(slog.String("op", op))

	var result dto.KaspiListCategoryResponse

	response, err := s.postgresStorage.ListKaspiCategory(ctx)
	if err != nil {
		return result, errorsShare.ErrInternalError.Error
	}

	err = copier.Copy(&result.Data, &response)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return result, err
	}

	return result, nil
}
