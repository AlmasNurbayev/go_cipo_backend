package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type NewsIDQueryRequest struct {
	Id int `form:"id" binding:"required" example:"5" validate:"min=0"`
}

type NewsIDItemResponse struct {
	Id             int       `json:"id"`
	Operation_date time.Time `json:"operation_date"`
	Title          string    `json:"title"`
	Data           string    `json:"data"`
	Image_path     string    `json:"image_path"`
	Changed_date   null.Time `json:"changed_date"`
}
