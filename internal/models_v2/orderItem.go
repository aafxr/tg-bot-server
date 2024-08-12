package modelsv2

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ProductId uint
	Quantity  int
	Pack      int
	OrderID   uint
}
