package modelsv2

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId string `gorm:"varchar(255)"`
	Quantity  int
	Pack      int
	OrderID   uint
}
