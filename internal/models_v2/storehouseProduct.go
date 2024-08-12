package modelsv2

import "gorm.io/gorm"

type StorehouseProduct struct {
	gorm.Model
	StorehouseID uint
	ProductID    uint
	Count        int
	Storehouse   Storehouse
	Product      Product
}
