package models

import "gorm.io/gorm"

type AppUser struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"column:id; type:int;unsigned;auto increment;not null;primaryKey;"`
	TgId         uint   `json:"tgID" gorm:"type:int;unsigned;not null;`
	Phone        string `josn:"phone" gorm:"column:phone;type:varchar(255)"`
	Country      string `josn:"country" gorm:"column:country;type:varchar(255)"`
	City         string `josn:"city" gorm:"column:city;type:varchar(255)"`
	TgUser       TgUser `json:"-"`
	Organization Organization
}
