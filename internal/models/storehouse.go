package models

import "gorm.io/gorm"

type Storehouse struct {
	gorm.Model
	ID      uint
	Country string
	City    string
	Address string
	Geo     string
}
