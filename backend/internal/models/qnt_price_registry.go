package models

import (
	"time"

	"github.com/guregu/null/v5"
)

type QntPriceRegistryEntity struct {
	Id int64 `json:"id" db:"id"`

	Sum              float32    `json:"sum" db:"sum"`
	Qnt              float32    `json:"qnt" db:"qnt"`
	Operation_date   time.Time  `json:"operation_date" db:"operation_date"`
	Discount_percent null.Float `json:"discount_percent" db:"discount_percent"`
	Discount_begin   null.Time  `json:"discount_begin" db:"discount_begin"`
	Discount_end     null.Time  `json:"discount_end" db:"discount_end"`

	Store_id         int64 `json:"store_id" db:"store_id"`
	Product_id       int64 `json:"product_id" db:"product_id"`
	Price_vid_id     int64 `json:"price_vid_id" db:"price_vid_id"`
	Size_id          int64 `json:"size_id" db:"size_id"`
	Registrator_id   int64 `json:"registrator_id" db:"registrator_id"`
	Product_group_id int64 `json:"product_group_id" db:"product_group_id"`
	Vid_modeli_id    int64 `json:"vid_modeli_id" db:"vid_modeli_id"`

	Size_name_1c        null.String `json:"size_name_1c" db:"size_name_1c"`
	Product_name        string      `json:"product_name" db:"product_name"`
	Product_create_date null.Time   `json:"product_created_at" db:"product_created_at"`

	Changed_date time.Time `json:"changed_date" db:"changed_date"`
	Create_date  time.Time `json:"create_date" db:"create_date"`
}
