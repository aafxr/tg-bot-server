package modelsv2

import "gorm.io/gorm"

type StorehouseProduct struct {
	gorm.Model
	StorehouseID uint
	ProductID    string `gorm:"varchar(255)"`
	Count        int
	Storehouse   Storehouse
	Product      Product
}
