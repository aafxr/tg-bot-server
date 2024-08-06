package models

import "gorm.io/gorm"

type ProductProperty struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"id;type:int;unsigned;auto increment;primaryKey;"`
	Name      string `json:"name" gorm:"name"`
	Value     string `json:"value" gorm:"value"`
	ProductID uint
}
