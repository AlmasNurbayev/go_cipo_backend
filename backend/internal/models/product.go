package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type ProductEntity struct {
	Id      int64  `json:"id" db:"id"`
	Id_1c   string `json:"id_1c" db:"id_1c"`
	Name_1c string `json:"name_1c" db:"name_1c"`
	Name    string `json:"name" db:"name"`

	Registrator_id   int64      `json:"registrator_id" db:"registrator_id"`
	Product_group_id int64      `json:"product_group_id" db:"product_group_id"`
	Product_vid_id   int64      `json:"product_vid_id" db:"product_vid_id"`
	Vid_modeli_id    null.Int64 `json:"vid_modeli_id" db:"vid_modeli_id"`
	Country_id       null.Int64 `json:"country_id" db:"country_id"`
	Brend_id         null.Int64 `json:"brend_id" db:"brend_id"`

	Artikul           string      `json:"artikul" db:"artikul"`
	Base_ed           string      `json:"base_ed" db:"base_ed"`
	Description       null.String `json:"description" db:"description"`
	Material_inside   null.String `json:"material_inside" db:"material_inside"`
	Material_podoshva null.String `json:"material_podoshva" db:"material_podoshva"`
	Material_up       null.String `json:"material_up" db:"material_up"`
	Sex               null.Int16  `json:"sex" db:"sex"`
	Product_folder    null.String `json:"product_folder" db:"product_folder"`
	Main_color        null.String `json:"main_color" db:"main_color"`
	Public_web        bool        `json:"public_web" db:"public_web"`

	Changed_date time.Time `json:"changed_date" db:"changed_date"`
	Create_date  time.Time `json:"create_date" db:"create_date"`
}

type ProductByIdEntity struct {
	Id      int64  `json:"id" db:"id"`
	Id_1c   string `json:"id_1c" db:"id_1c"`
	Name_1c string `json:"name_1c" db:"name_1c"`
	Name    string `json:"name" db:"name"`

	Registrator_id   int64      `json:"registrator_id" db:"registrator_id"`
	Product_group_id int64      `json:"product_group_id" db:"product_group_id"`
	Product_vid_id   int64      `json:"product_vid_id" db:"product_vid_id"`
	Vid_modeli_id    null.Int64 `json:"vid_modeli_id" db:"vid_modeli_id"`
	Country_id       null.Int64 `json:"country_id" db:"country_id"`
	Brend_id         null.Int64 `json:"brend_id" db:"brend_id"`

	Artikul           string      `json:"artikul" db:"artikul"`
	Base_ed           string      `json:"base_ed" db:"base_ed"`
	Description       null.String `json:"description" db:"description"`
	Material_inside   null.String `json:"material_inside" db:"material_inside"`
	Material_podoshva null.String `json:"material_podoshva" db:"material_podoshva"`
	Material_up       null.String `json:"material_up" db:"material_up"`
	Sex               null.Int16  `json:"sex" db:"sex"`
	Product_folder    null.String `json:"product_folder" db:"product_folder"`
	Main_color        null.String `json:"main_color" db:"main_color"`
	Public_web        bool        `json:"public_web" db:"public_web"`

	Changed_date time.Time `json:"changed_date" db:"changed_date"`
	Create_date  time.Time `json:"create_date" db:"create_date"`

	Vid_modeli    VidModeliEntity     `json:"vid_modeli" db:"vid_modeli"`
	Product_group ProductsGroupEntity `json:"product_group" db:"product_group"`
}
