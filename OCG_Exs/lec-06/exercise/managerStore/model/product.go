package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	OrderDetails []OrderDetail
}

var Products []Product
