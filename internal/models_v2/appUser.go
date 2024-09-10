package modelsv2

import "gorm.io/gorm"

type AppUser struct {
	gorm.Model
	ID            uint           `json:"id" gorm:"primatyKey;unsigned;autoIncrement"`
	Phone         int            `json:"phone" gorm:"column:phone"`
	Country       string         `json:"country" gorm:"column:country;type:varchar(255)"`
	City          string         `json:"city" gorm:"column:city;type:varchar(255)"`
	TgUser        TgUser         `json:"tgUser,omitempty" gorm:"foreighKey:ID;refference:TgUserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	StorehouseId  string         `json:"storehouseId" gorm:"column:storehouse_id;type:varchar(40)"`
	Organizations []Organization `json:"organizations,omitempty;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Orders        []Order        `json:"orders,omitempty;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
