package modelsv2

import "gorm.io/gorm"

type AppUser struct {
	gorm.Model
	FirstName     string         `gorm:"varchar(255)"`
	LastName      string         `gorm:"varchar(255)"`
	Phone         string         `josn:"phone" gorm:"column:phone;type:varchar(255)"`
	Country       string         `josn:"country" gorm:"column:country;type:varchar(255)"`
	City          string         `josn:"city" gorm:"column:city;type:varchar(255)"`
	TgUserID      uint           `josn:"tgUserId"`
	TgUser        TgUser         `json:"tgUser,omitempty" gorm:"foreighKey:ID;refference:TgUserID"`
	Organizations []Organization `json:"organizations,omitempty"`
	Orders        []Order        `json:"orders,omitempty"`
}
