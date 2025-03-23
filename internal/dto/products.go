package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type ProductsQueryRequest struct {
	Search_name   string  `form:"search_name" binding:"omitempty" example:"name" query:"search_name" validate:"omitempty,gte=3"`
	MinPrice      float32 `form:"min_price" binding:"omitempty" example:"100" query:"minPrice" validate:"omitempty,min=0"`
	MaxPrice      float32 `form:"max_price" binding:"omitempty" example:"1000" query:"maxPrice" validate:"omitempty,min=0"`
	Sort          string  `form:"sort" binding:"omitempty" example:"sum-desc" query:"Sort"`
	Size          []int64 `form:"size" binding:"omitempty" example:"31" query:"size"`
	Vid_modeli    []int64 `form:"vid_modeli" binding:"omitempty" example:"1" query:"vid_modeli"`
	Product_group []int64 `form:"product_group" binding:"omitempty" example:"1" query:"product_group" `
	Take          int     `form:"page" binding:"omitempty" example:"20" query:"page" validate:"omitempty,min=0,max=100"`
	Skip          int     `form:"skip" binding:"omitempty" example:"0" query:"skip" validate:"omitempty,min=0"`
}

type ProductsResponse struct {
	Data          []ProductsItemResponse `json:"data"`
	Full_count    int                    `json:"full_count"`
	Current_count int                    `json:"current_count"`
}

type ProductsItemResponse struct {
	Product_id          int64           `json:"product_id"`
	Product_create_date time.Time       `json:"product_create_date"`
	Sum                 float32         `json:"sum"`
	Product_group_id    int64           `json:"product_group_id"`
	Product_name        string          `json:"product_name"`
	Qnt_price           []qnt_price     `json:"qnt_price"`
	Artikul             string          `json:"artikul"`
	Name                string          `json:"name"`
	Description         null.String     `json:"description"`
	Material_podoshva   null.String     `json:"material_podoshva"`
	Material_up         null.String     `json:"material_up"`
	Material_inside     null.String     `json:"material_inside"`
	Sex                 null.String     `json:"sex"`
	Product_group_name  string          `json:"product_group_name"`
	Vid_modeli_name     null.String     `json:"vid_modeli_name"`
	Vid_modeli_id       int             `json:"vid_modeli_id"`
	Image_registry      []imageRegistry `json:"image_registry"`
	Image_active_path   string          `json:"image_active_path"`
	Create_date         time.Time       `json:"create_date"`
}

type qnt_price struct {
	Size     string  `json:"size"`
	Sum      float32 `json:"sum"`
	Qnt      float32 `json:"qnt"`
	Store_id []int   `json:"store_id"`
}

type imageRegistry struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	Main      bool   `json:"main"`
	Full_name string `json:"full_name"`
}
