package models

import "time"

type KaspiCategoriesEntity struct {
	Id           int64    `db:"id"`
	Name_kaspi   string   `db:"name_kaspi"`
	Title_kaspi  string   `db:"title_kaspi"`
	First_size   string   `db:"first_size"`
	Last_size    string   `db:"last_size"`
	Size_kaspi   []string `db:"size_kaspi"`
	Gender_kaspi []string `db:"gender_kaspi"`
	Model_kaspi  []string `db:"model_kaspi"`

	Material_kaspi []string `db:"material_kaspi"`
	Season_kaspi   []string `db:"season_kaspi"`
	Colour_kaspi   []string `db:"colour_kaspi"`

	Attributes_list []map[string]any `json:"attributes_list"`

	Changed_date time.Time `db:"changed_date"`
	Create_date  time.Time `db:"create_date"`
}
