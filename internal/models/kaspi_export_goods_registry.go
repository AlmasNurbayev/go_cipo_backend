package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type KaspiExportGoodsRegistryEntity struct {
	Id                  int64       `db:"id"`
	KaspiOrganizationId int64       `db:"kaspi_organization_id"`
	ProductId           int64       `db:"product_id"`
	SendedBody          string      `db:"sended_body"`
	SendedCategory      string      `db:"sended_category"`
	SendedStatus        int         `db:"sended_status"`
	ResponseId          null.String `db:"response_id"`
	ResponseStatus      null.String `db:"response_status"`
	Errors              []string    `db:"errors"`
	CreateDate          time.Time   `db:"create_date"`
	ChangedDate         time.Time   `db:"changed_date"`
}
