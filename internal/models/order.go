package models

type Order struct {
	ID             uint   `json:"id" gorm:"column:id;type:int;unsigned;auto increment;primaryKey;"`
	Status         string `json:"status" gorm:"column:status;type:varchar(255);"`
	AppUserID      uint
	OrganizationID uint
}
