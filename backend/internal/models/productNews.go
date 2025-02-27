package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type ProductNewsEntity struct {
	Product_id          int64     `json:"product_id" db:"product_id"`
	Product_name        string    `json:"product_name"  db:"product_name"`
	Sum                 float32   `json:"sum" db:"sum"`
	Product_create_date time.Time `json:"product_create_date" db:"product_create_date"`
	//	Qnt_price           qnt_price               `json:"qnt_price" db:"qnt_price"`
	Name               string                `json:"name" db:"name"`
	Artikul            string                `json:"artikul" db:"artikul"`
	Description        null.String           `json:"description" db:"description"`
	Material_up        null.String           `json:"material_up" db:"material_up"`
	Material_inside    null.String           `json:"material_inside" db:"material_inside"`
	Material_podoshva  null.String           `json:"material_podoshva" db:"material_podoshva"`
	Sex                null.Int16            `json:"sex" db:"sex"`
	Product_group_name string                `json:"product_group_name" db:"product_group_name"`
	Vid_modeli_name    string                `json:"vid_modeli_name" db:"vid_modeli_name"`
	Vid_modeli_id      null.Int64            `json:"vid_modeli_id" db:"vid_modeli_id"`
	Image_registry     []imageRegistryEntity `json:"image_registry" db:"image_registry"`
	Image_active_path  string                `json:"image_active_path" db:"image_active_path"`
}

type imageRegistryEntity struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Active    bool   `json:"active" db:"active"`
	Main      bool   `json:"main" db:"main"`
	Full_name string `json:"full_name" db:"full_name"`
}
