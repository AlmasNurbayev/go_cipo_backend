package dto

import "github.com/guregu/null/v5"

type ProductNewsQueryRequest struct {
	News int `form:"news" binding:"omitempty" example:"5" validate:"min=0,max=11"`
}

type ProductNewsResponse struct {
	Product_id          int64       `json:"product_id"`
	Product_name        string      `json:"product_name"`
	Sum                 float32     `json:"sum"`
	Product_create_date string      `json:"product_create_date"`
	Qnt_price           []qnt_price `json:"qnt_price"`
	Name                string      `json:"name"`
	Artikul             string      `json:"artikul"`
	Description         null.String `json:"description"`
	Material_up         null.String `json:"material_up"`
	Material_inside     null.String `json:"material_inside"`
	Material_podoshva   null.String `json:"material_podoshva"`
	Sex                 null.Int16  `json:"sex"`
	Product_group_name  string      `json:"product_group_name"`
	Vid_modeli_name     string      `json:"vid_modeli_name"`
	Vid_modeli_id       null.Int64  `json:"vid_modeli_id"`
	Image_registry      []imR       `json:"image_registry"`
	Image_active_path   string      `json:"image_active_path"`
}

type imR struct {
	Id        int64  `json:"id"`
	Full_name string `json:"full_name"`
	Name      string `json:"name"`
	Main      bool   `json:"main"`
	Active    bool   `json:"active"`
}
