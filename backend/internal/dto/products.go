package dto

type ProductsQueryRequest struct {
	Search_name   string `form:"search_name" binding:"omitempty" example:"name"`
	MinPrice      string `form:"min_price" binding:"omitempty" example:"100"`
	MaxPrice      string `form:"max_price" binding:"omitempty" example:"1000"`
	Sort          string `form:"sort" binding:"omitempty" example:"sum-desc"`
	Size          string `form:"size" binding:"omitempty" example:"31"`
	Vid_modeli    string `form:"vid_modeli" binding:"omitempty" example:"1"`
	Product_group string `form:"product_group" binding:"omitempty" example:"1"`
	Take          string `form:"page" binding:"omitempty" example:"20"`
	Skip          string `form:"skip" binding:"omitempty" example:"0"`
}

type ProductsResponse struct {
	Data          []ProductsItemResponse `json:"data"`
	Full_count    int                    `json:"full_count"`
	Current_count int                    `json:"current_count"`
}

type ProductsItemResponse struct {
	Product_id          int64            `json:"product_id"`
	Product_create_date string           `json:"product_create_date"`
	Sum                 float32          `json:"sum"`
	Product_group_id    int64            `json:"product_group_id"`
	Product_name        string           `json:"product_name"`
	Qnt_price           *[]qnt_price     `json:"qnt_price"`
	Artikul             string           `json:"artikul"`
	Name                string           `json:"name"`
	Description         *string          `json:"description"`
	Material_podoshva   *string          `json:"material_podoshva"`
	Material_up         *string          `json:"material_up"`
	Material_inside     *string          `json:"material_inside"`
	Sex                 *string          `json:"sex"`
	Product_group_name  string           `json:"product_group_name"`
	Vid_modeli_name     *string          `json:"vid_modeli_name"`
	Vid_modeli_id       int              `json:"vid_modeli_id"`
	Image_registry      *[]imageRegistry `json:"image_registry"`
	Image_active_path   string           `json:"image_active_path"`
	Create_date         string           `json:"create_date"`
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
