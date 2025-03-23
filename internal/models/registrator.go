package models

import (
	"time"
)

type RegistratorEntity struct {
	Id             int64     `json:"id" db:"id"`
	Operation_date time.Time `json:"operation_date" db:"operation_date"`
	Name_folder    string    `json:"name_folder" db:"name_folder"`
	Name_file      string    `json:"name_file" db:"name_file"`
	User_id        int64     `json:"user_id" db:"user_id"`
	Date_schema    time.Time `json:"date_schema" db:"date_schema"`
	Id_catalog     string    `json:"id_catalog" db:"id_catalog"`
	Id_class       string    `json:"id_class" db:"id_class"`
	Name_catalog   string    `json:"name_catalog" db:"name_catalog"`
	Name_class     string    `json:"name_class" db:"name_class"`
	Ver_schema     string    `json:"ver_schema" db:"ver_schema"`
	Only_change    bool      `json:"only_change" db:"only_change"`
	Changed_date   time.Time `json:"changed_date" db:"changed_date"`
	Create_date    time.Time `json:"create_date" db:"create_date"`
}
