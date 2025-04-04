package dto

import "time"

type KaspiAddCategoryRequest struct {
	Name_kaspi     string `example:"Master - Girl boots" validate:"required,gte=3" json:"name_kaspi"`
	OrganizationId int64  `example:"1" validate:"required,gte=1" json:"organization_id"`
}

type KaspiAddCategoryResponse struct {
	Id           int64    `json:"id"`
	Name_kaspi   string   `json:"name_kaspi"`
	Title_kaspi  string   `json:"title_kaspi"`
	First_size   string   `json:"first_size"`
	Last_size    string   `json:"last_size"`
	Size_kaspi   []string `json:"size_kaspi"`
	Gender_kaspi []string `json:"gender_kaspi"`
	Model_kaspi  []string `json:"model_kaspi"`

	Material_kaspi []string `json:"material_kaspi"`
	Season_kaspi   []string `json:"season_kaspi"`
	Colour_kaspi   []string `json:"colour_kaspi"`

	Attributes_list []map[string]any `json:"attributes_list"`

	Changed_date time.Time `json:"changed_date"`
	Create_date  time.Time `json:"create_date"`
}

type KaspiListCategoryResponse struct {
	Data []KaspiAddCategoryResponse `json:"data"`
}
