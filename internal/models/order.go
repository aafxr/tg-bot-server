package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID             uint
	Status         string `json:"status" gorm:"column:status;type:varchar(255);"`
	AppUserID      uint
	OrganizationID uint
}
