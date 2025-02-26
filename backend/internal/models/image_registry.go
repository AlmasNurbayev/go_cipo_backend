package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type ImageRegistryEntity struct {
	Id               int64       `json:"id" db:"id"`
	Main             bool        `json:"main" db:"main"`
	Main_change_date time.Time   `json:"main_change_date" db:"main_change_date"`
	Resolution       null.String `json:"resolution" db:"resolution"`
	Size             int         `json:"size" db:"size"`
	Full_name        string      `json:"full_name" db:"full_name"`
	Name             string      `json:"name" db:"name"`
	Path             string      `json:"path" db:"path"`

	Operation_date     time.Time `json:"operation_date" db:"operation_date"`
	Active             bool      `json:"active" db:"active"`
	Active_change_date time.Time `json:"active_change_date" db:"active_change_date"`

	Registrator_id int64 `json:"registrator_id" db:"registrator_id"`
	Product_id     int64 `json:"product_id" db:"product_id"`

	Changed_date time.Time `json:"changed_date" db:"changed_date"`
	Create_date  time.Time `json:"create_date" db:"create_date"`
}
