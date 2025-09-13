package dto

type ProductsFilterResponse struct {
	Size          []idName1c `json:"size"`
	Product_group []idName1c `json:"product_group"`
	Vid_modeli    []idName1c `json:"vid_modeli"`
	Brend         []idName1c `json:"brend"`
	Store         []idName1c `json:"store"`
	Nom_vid       []string   `json:"nom_vid"`
}

type idName1c struct {
	Id      int    `json:"id"`
	Name_1c string `json:"name_1c"`
}
