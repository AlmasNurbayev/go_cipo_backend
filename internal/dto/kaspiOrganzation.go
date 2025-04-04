package dto

import "time"

type KaspiAddOrganizationRequest struct {
	Name            string `example:"ТОО Рога" validate:"required,gte=3" json:"name"`
	Kaspi_id        string `example:"123456" validate:"required,gte=1" json:"kaspi_id"`
	Kaspi_api_token string `example:"123456" validate:"required,gte=1" json:"kaspi_api_token"`
}

type KaspiAddOrganizationResponse struct {
	Id           int64     `json:"id"`
	Kaspi_id     string    `json:"kaspi_id"`
	Name         string    `json:"name"`
	Changed_date time.Time `json:"changed_date"`
	Create_date  time.Time `json:"create_date"`
}

type KaspiListOrganizationResponse struct {
	Data []KaspiAddOrganizationResponse `json:"data"`
}
