package models

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

var Products []Product
