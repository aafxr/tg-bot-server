package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"column:id;type:int;unsigned;auto increment;not null;primaryKey;"`
	Name      string `json:"name" gorm:"column:name;type:varchar(255)"`
	Country   string `json:"country" gorm:"column:country;type:varchar(255)"`
	City      string `json:"city" gorm:"column:city;type:varchar(255)"`
	INN       uint   `json:"inn" gorm:"column:inn;type:int;unsigned"`
	AppUserID uint
}
