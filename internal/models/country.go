package models

import (
	"time"
)

type CountryEntity struct {
	Id             int64     `json:"id" db:"id"`
	Id_1c          string    `json:"id_1c" db:"id_1c"`
	Name_1c        string    `json:"name_1c" db:"name_1c"`
	Registrator_id int64     `json:"registrator_id" db:"registrator_id"`
	Changed_date   time.Time `json:"changed_date" db:"changed_date"`
	Create_date    time.Time `json:"create_date" db:"create_date"`
}
