package modelsv2

import "gorm.io/gorm"

type Client struct {
	Phone     string
	AppUserID uint
	FirstName string
	LastName  string
}

type Order struct {
	gorm.Model
	AppUserID      uint
	Status         string `json:"status" gorm:"column:status;type:varchar(255);"`
	Comment        string
	OrderItems     []OrderItem
	OrganizationID uint
	StorehouseID   uint
	Storehouse     Storehouse
}
