package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type NewsEntity struct {
	Id             int64       `json:"id" db:"id"`
	Title          string      `json:"title" db:"title"`
	Data           string      `json:"data" db:"data"`
	Image_path     null.String `json:"image_path" db:"image_path"`
	Operation_date time.Time   `json:"operation_date" db:"operation_date"`
	Changed_date   time.Time   `json:"changed_date" db:"changed_date"`
	Create_date    time.Time   `json:"create_date" db:"create_date"`
}
