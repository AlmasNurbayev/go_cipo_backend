package models

import (
	"time"
)

type KaspiOrganizationEntity struct {
	Id                   int64  `db:"id"`
	Name                 string `db:"name"`
	Kaspi_id             string `db:"kaspi_id"`
	Kaspi_api_token_hash string `db:"kaspi_api_token_hash"`
	Is_active            bool   `db:"is_active"`

	Changed_date time.Time `db:"changed_date"`
	Create_date  time.Time `db:"create_date"`
}
