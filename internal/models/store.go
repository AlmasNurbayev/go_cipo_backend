package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type StoreEntity struct {
	Id_1c                string        `db:"id_1c"`
	Id                   int64         `db:"id"`
	Name_1c              string        `db:"name_1c"`
	Address              null.String   `db:"address"`
	Link_2gis            null.String   `db:"link_2gis"`
	Phone                null.String   `db:"phone"`
	City                 null.String   `db:"city"`
	Image_path           null.String   `db:"image_path"`
	Public               bool          `db:"public"`
	Working_hours        working_hours `db:"working_hours"`
	Yandex_widget_url    string        `db:"yandex_widget_url"`
	Doublegis_widget_url string        `db:"doublegis_widget_url"`
	Registrator_id       int64         `db:"registrator_id"`
	Store_kaspi_id       null.String   `db:"store_kaspi_id"`
	Changed_date         time.Time     `json:"changed_date" db:"changed_date"`
	Create_date          time.Time     `json:"create_date" db:"create_date"`
}

type working_hours struct {
	D01 string `db:"d01"`
	D02 string `db:"d02"`
	D03 string `db:"d03"`
	D04 string `db:"d04"`
	D05 string `db:"d05"`
	D06 string `db:"d06"`
	D07 string `db:"d07"`
}
