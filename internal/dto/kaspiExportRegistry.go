package dto

import (
	"time"

	"github.com/guregu/null/v5"
)

type KaspiExportGoodsRegistryItem struct {
	Id                  int64       `json:"id"`
	KaspiOrganizationId int64       `json:"kaspi_organization_id"`
	ProductId           int64       `json:"product_id"`
	SendedBody          string      `json:"sended_body"`
	SendedCategory      string      `json:"sended_category"`
	SendedStatus        int         `json:"sended_status"`
	ResponseId          null.String `json:"response_id"`
	ResponseStatus      null.String `json:"response_status"`
	CreateDate          time.Time   `json:"create_at"`
	ChangedDate         time.Time   `json:"changed_at"`
}
