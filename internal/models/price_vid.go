package models

import (
	"time"
)

type PriceVidEntity struct {
	Id      int64  `json:"id" db:"id"`
	Name_1c string `json:"name_1c" db:"name_1c"`
	Id_1c   string `json:"id_1c" db:"id_1c"`

	Active             bool      `json:"active" db:"active"`
	Active_change_date time.Time `json:"active_change_date" db:"active_change_date"`

	Registrator_id int64 `json:"registrator_id" db:"registrator_id"`

	Changed_date time.Time `json:"changed_date" db:"changed_date"`
	Create_date  time.Time `json:"create_date" db:"create_date"`
}
