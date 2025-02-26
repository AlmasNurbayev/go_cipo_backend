package dto

type NewsQueryRequest struct {
	News *int `form:"news" binding:"required" example:"5"`
}

type NewsItemResponse struct {
	Id             int    `json:"id"`
	Operation_date string `json:"operation_date"`
	Title          string `json:"title"`
	Data           string `json:"data"`
	Image_path     string `json:"image_path"`
	Changed_date   string `json:"changed_date"`
}
