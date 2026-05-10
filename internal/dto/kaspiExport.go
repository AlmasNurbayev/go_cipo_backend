package dto

type ExportProductRequest struct {
	OrganizationId int64 `json:"organization_id"`
	Data           []struct {
		SKU         string `json:"sku"`
		Title       string `json:"title"`
		Brand       string `json:"brand"`
		Category    string `json:"category"`
		Description string `json:"description"`
		FamilyID    string `json:"familyId"`
		Attributes  []struct {
			Code  string `json:"code"`
			Value any    `json:"value"`
		} `json:"attributes"`
		Images []struct {
			Url string `json:"url"`
		} `json:"images"`
	} `json:"data"`
}

// type DynamicValue struct {
// 	String      string
// 	StringSlice []string
// 	Bool        bool
// 	Number      float64
// 	Type        string // Чтобы знать, какой тип пришел: "string", "slice", "bool", "number"
// }

// в любом варианте приходит 200
type ExportProductResponse struct {
	Code   string `json:"code"` //
	Status string `json:"status"`
	Errors []struct {
		Code   int    `json:"code"`
		Detail string `json:"detail"`
		Status string `json:"status"`
		Title  string `json:"title"`
	} `json:"errors"`
}
