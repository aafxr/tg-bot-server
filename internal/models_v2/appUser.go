package modelsv2

import "gorm.io/gorm"

type AppUser struct {
	gorm.Model
	ID            uint           `josn:"id" gorm:"primatyKey;unsigned;autoIncrement"`
	FirstName     string         `gorm:"varchar(255)"`
	LastName      string         `gorm:"varchar(255)"`
	Phone         string         `josn:"phone" gorm:"column:phone;type:varchar(255)"`
	Country       string         `josn:"country" gorm:"column:country;type:varchar(255)"`
	City          string         `josn:"city" gorm:"column:city;type:varchar(255)"`
	TgUser        TgUser         `json:"tgUser,omitempty" gorm:"foreighKey:ID;refference:TgUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Organizations []Organization `json:"organizations,omitempty;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Orders        []Order        `json:"orders,omitempty;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
