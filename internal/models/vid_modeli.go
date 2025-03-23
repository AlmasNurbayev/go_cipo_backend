package models

import (
	"time"
)

type VidModeliEntity struct {
	Id             int64  `db:"id"`
	Id_1c          string `db:"id_1c"`
	Name_1c        string `db:"name_1c"`
	Registrator_id int64  `db:"registrator_id"`

	Changed_date time.Time `json:"changed_date" db:"changed_date"`
	Create_date  time.Time `json:"create_date" db:"create_date"`
}
