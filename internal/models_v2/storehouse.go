package modelsv2

import "gorm.io/gorm"

type Storehouse struct {
	gorm.Model
	ID      uint
	Country string
	City    string
	Address string
	Geo     string
}
