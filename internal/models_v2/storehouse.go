package modelsv2

import "gorm.io/gorm"

type Storehouse struct {
	gorm.Model
	ID      uint
	Country string `gorm:"varchar(255)"`
	City    string `gorm:"varchar(255)"`
	Address string `gorm:"varchar(255)"`
	Geo     string `gorm:"varchar(255)"`
}
