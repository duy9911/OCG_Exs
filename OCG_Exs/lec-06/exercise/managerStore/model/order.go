package model

import "gorm.io/gorm"

//many to many relationship, order can have many product and product too
type Order struct {
	gorm.Model
	CustomerID    uint   `json:"customer_id"`
	OrderDate     string `json:"order_date"`
	PaymentMethod string `json:"payment_method"`
	OrderDetails  []OrderDetail
}

var Orders []Order

type OrderDetail struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	OrderID   uint    `json:"order_id"`
	Quantity  float64 `json:"quantity"`
}

var OrderDetails []OrderDetail
