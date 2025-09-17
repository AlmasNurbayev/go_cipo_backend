package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type ProductNewsEntity struct {
	Product_id          int64             `json:"product_id" db:"product_id"`
	Product_name        string            `json:"product_name"  db:"product_name"`
	Product_create_date time.Time         `json:"product_create_date" db:"product_create_date"`
	Name                string            `json:"name" db:"name"`
	Artikul             string            `json:"artikul" db:"artikul"`
	Description         null.String       `json:"description" db:"description"`
	Material_up         null.String       `json:"material_up" db:"material_up"`
	Material_inside     null.String       `json:"material_inside" db:"material_inside"`
	Material_podoshva   null.String       `json:"material_podoshva" db:"material_podoshva"`
	Sex                 null.Int16        `json:"sex" db:"sex"`
	Product_group_name  string            `json:"product_group_name" db:"product_group_name"`
	Vid_modeli_name     string            `json:"vid_modeli_name" db:"vid_modeli_name"`
	Vid_modeli_id       null.Int64        `json:"vid_modeli_id" db:"vid_modeli_id"`
	Nom_vid             null.String       `json:"nom_vid" db:"nom_vid"`
	Qnt_price           []qnt_price_short `json:"qnt_price" db:"qnt_price"`
	Image_active_path   string            `json:"image_active_path" db:"image_active_path"`
}

type qnt_price_short struct {
	Size string  `json:"size" db:"size"`
	Sum  float32 `json:"sum" db:"sum"`
}
