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
	op := "services.KaspiAddCategory"
	log := s.log.With(slog.String("op", op))

	newCategory := models.KaspiCategoriesEntity{
		Name_kaspi: data.Name_kaspi,
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
			log.Debug("selected organization", slog.String("organization_name", organization.Name))
		}
	}
	if token == "" {
		// если что-то не так с токеном
		return response, errorsShare.ErrInternalError.Error
	}

	// получаем список категорий и проверяем на соответствие кода
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
		//log.Error("Error", slog.String("err", errCategories.Error()))
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
	// str, _ := utils.PrintAsJSON(attributes)
	// fmt.Println("attributes", string(*str))

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

	result, err := s.postgresStorage.CreateKaspiCategory(ctx, newCategory)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
		return response, err
	}

	err = copier.Copy(&response, &result)
	if err != nil {
		log.Error("", slog.String("err", err.Error()))
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
