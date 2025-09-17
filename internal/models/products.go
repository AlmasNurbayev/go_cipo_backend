package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type ProductsItemEntity struct {
	Product_id          int64     `json:"product_id" db:"product_id"`
	Product_name        string    `json:"product_name"  db:"product_name"`
	Sum                 float32   `json:"sum" db:"sum"`
	Product_create_date time.Time `json:"product_create_date" db:"product_create_date"`

	Name               string                `json:"name" db:"name"`
	Artikul            string                `json:"artikul" db:"artikul"`
	Description        null.String           `json:"description" db:"description"`
	Material_up        null.String           `json:"material_up" db:"material_up"`
	Material_inside    null.String           `json:"material_inside" db:"material_inside"`
	Material_podoshva  null.String           `json:"material_podoshva" db:"material_podoshva"`
	Sex                null.Int              `json:"sex" db:"sex"`
	Product_group_name string                `json:"product_group_name" db:"product_group_name"`
	Product_group_id   int64                 `json:"product_group_id" db:"product_group_id"`
	Vid_modeli_name    string                `json:"vid_modeli_name" db:"vid_modeli_name"`
	Vid_modeli_id      null.Int64            `json:"vid_modeli_id" db:"vid_modeli_id"`
	Image_registry     []imageRegistryEntity `json:"image_registry" db:"image_registry"`
	Qnt_price          []qnt_price           `json:"qnt_price" db:"qnt_price"`
	Image_active_path  string                `json:"image_active_path" db:"image_active_path"`
	Nom_vid            null.String           `json:"nom_vid" db:"nom_vid"`

	Create_date time.Time `db:"create_date"`

	Kaspi_in_sale  bool        `db:"kaspi_in_sale"`
	Kaspi_category null.String `db:"kaspi_category"`
	Main_color     null.String `db:"main_color"`
}

type imageRegistryEntity struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Active    bool   `json:"active" db:"active"`
	Main      bool   `json:"main" db:"main"`
	Full_name string `json:"full_name" db:"full_name"`
}

type qnt_price struct {
	Size     string  `json:"size" db:"size"`
	Sum      float32 `json:"sum" db:"sum"`
	Qnt      float32 `json:"qnt" db:"qnt"`
	Store_id []int64 `json:"store_id" db:"store_id"`
}
