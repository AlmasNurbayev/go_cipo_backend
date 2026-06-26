package dto

type ProductsFilterResponse struct {
	Size          []IdName1c `json:"size"`
	Product_group []IdName1c `json:"product_group"`
	Vid_modeli    []IdName1c `json:"vid_modeli"`
	Brend         []IdName1c `json:"brend"`
	Store         []IdName1c `json:"store"`
	Nom_vid       []string   `json:"nom_vid"`
}

type IdName1c struct {
	Id      int    `json:"id"`
	Name_1c string `json:"name_1c"`
}
