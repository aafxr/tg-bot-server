package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Count     uint    `json:"count" gorm:"count"`
	Price     float32 `json:"price" gorm:"price"`
	Currency  string  `json:"currency" gorm:"currency"`
	OrderID   uint
	ProductID uint
}
