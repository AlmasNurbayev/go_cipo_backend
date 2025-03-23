package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type StoresResponse struct {
	Stores []StoreResponse `json:"stores"`
}

type StoreResponse struct {
	Id                   int           `json:"id"`
	Id_1c                string        `json:"id_1c"`
	Name_1c              string        `json:"name_1c"`
	Address              string        `json:"address"`
	Link_2gis            string        `json:"link_2gis"`
	Phone                string        `json:"phone"`
	City                 string        `json:"city"`
	Image_path           string        `json:"image_path"`
	Public               bool          `json:"public"`
	Working_hours        Working_hours `json:"working_hours"`
	Yandex_widget_url    string        `json:"yandex_widget_url"`
	Doublegis_widget_url string        `json:"doublegis_widget_url"`
	Store_kaspi_id       string        `json:"store_kaspi_id"`
	Registrator_id       int           `json:"registrator_id"`
	Create_date          time.Time     `json:"create_date"`
	Changed_date         null.Time     `json:"changed_date"`
}

type Working_hours struct {
	D01 string `json:"d01"`
	D02 string `json:"d02"`
	D03 string `json:"d03"`
	D04 string `json:"d04"`
	D05 string `json:"d05"`
	D06 string `json:"d06"`
	D07 string `json:"d07"`
}
