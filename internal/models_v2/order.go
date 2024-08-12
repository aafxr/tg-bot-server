package modelsv2

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status         string `json:"status" gorm:"column:status;type:varchar(255);"`
	Comment        string
	OrderItems     []OrderItem
	TgUserID       uint
	OrganizationID uint
	StorehouseID   uint
	Storehouse     Storehouse
}
