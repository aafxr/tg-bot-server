package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID         uint    `json:"id" gorm:"column:id;type:int;unsigned;auto increment;not null;primaryKey"`
	Title      string  `json:"title" gorm:"column:title;"`
	Currency   string  `json:"currency" gorm:"column:currency;"`
	Price      float32 `json:"price" gorm:"column:price;"`
	Preview    string  `json:"preview" gorm:"column:preview;"`
	Properties []ProductProperty
	OrderItems []OrderItem
}
