package modelsv2

import (
	"errors"

	"gorm.io/gorm"
)

type Client struct {
	Phone     string
	AppUserID uint
	FirstName string
	LastName  string
}

type Order struct {
	gorm.Model
	ID             uint        `json:"id" gorm:"primatyKey;unsigned;autoIncrement"`
	AppUserID      uint        `json:"userID"`
	Status         string      `json:"status" gorm:"column:status;type:varchar(255);"`
	Comment        string      `json:"comment"`
	OrderItems     []OrderItem `json:"items" gorm:"type:json;serializer:json"`
	OrganizationID uint        `json:"companyID"`
	// StorehouseID   uint        `json:"storehouseID"`
	// Storehouse     Storehouse
}

func (o Order) Validate() (bool, error) {
	if o.AppUserID == 0 {
		return false, errors.New("отсутствует id пользователя")
	}
	if len(o.OrderItems) == 0 {
		return false, errors.New("пустой заказ")
	}
	for _, oi := range o.OrderItems {
		res, err := oi.Validate()
		if err != nil {
			return res, err
		}
	}
	return true, nil
}
