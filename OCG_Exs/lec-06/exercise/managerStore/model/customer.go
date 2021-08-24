package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	CustomerInfo string `json:"customer_info"`
	Orders       []Order
}

var Customers []Customer
