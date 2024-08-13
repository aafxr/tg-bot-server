package modelsv2

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name;type:varchar(255)"`
	FullName  string `json:"fullName" gorm:"column:full_name;type:varchar(255)"`
	Address   string `json:"address" gorm:"column:address;type:varchar(255)"`
	Country   string `json:"country" gorm:"column:country;type:varchar(255)"`
	City      string `json:"city" gorm:"column:city;type:varchar(255)"`
	INN       uint   `json:"inn" gorm:"column:inn;type:int;unsigned"`
	AppUserID uint
}
