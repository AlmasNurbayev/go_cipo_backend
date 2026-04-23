package dto

import "time"

type StockJSON struct {
	ExportDate  time.Time `json:"ExportDate"`
	ExportPath  string    `json:"ExportPath"`
	BasePrefix  string    `json:"BasePrefix"`
	NameLog     string    `json:"NameLog"`
	TotalQnt    int64     `json:"TotalQnt"`
	DurationSec int64     `json:"DurationSec"`
	Qnt         []struct {
		NomCode              string               `json:"NomCode"`
		NomGUID              string               `json:"NomGUID"`
		NomName              string               `json:"NomName"`
		NomArticle           string               `json:"NomArticle"`
		TovarGroupName       string               `json:"TovarGroupName"`
		TovarGroupGUID       string               `json:"TovarGroupGUID"`
		EdIzmName            string               `json:"EdIzmName"`
		EdIzmGUID            string               `json:"EdIzmGUID"`
		MarkaName            string               `json:"MarkaName"`
		MarkaGUID            string               `json:"MarkaGUID"`
		VidName              string               `json:"VidName"`
		VidGUID              string               `json:"VidGUID"`
		AdditionalProperties []AdditionalProperty `json:"AdditionalProperties"`
		Images               []string             `json:"Images"`
		Stocks               []Stock              `json:"Stocks"`
	}
}

type AdditionalProperty []struct {
	NameProperty     string `json:"NameProperty"`
	GUIDProperty     string `json:"GUIDProperty"`
	StrValueProperty string `json:"StrValueProperty"`
}

type Stock struct {
	Warehouse     string  `json:"Warehouse"`
	WarehouseGUID string  `json:"WarehouseGUID"`
	Char          string  `json:"Char"`
	CharGUID      string  `json:"CharGUID"`
	Quantity      int64   `json:"Quantity"`
	Price         float32 `json:"Price"`
	Barcode       string  `json:"Barcode"`
}
