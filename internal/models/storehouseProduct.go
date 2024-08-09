package models

import "gorm.io/gorm"

type StorehouseProduct struct {
	gorm.Model
	StorehouseID uint
	ProductID    uint
	count        int
	Storehouse   Storehouse
	Product      Product
}
