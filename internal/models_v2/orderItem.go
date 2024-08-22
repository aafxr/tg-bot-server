package modelsv2

import (
	"errors"
	"fmt"
)

type OrderItem struct {
	ID               string  `json:"id" gorm:"type:varchar(10)"`
	Title            string  `json:"title" gorm:"type:varchar(255)"`
	Count            float32 `json:"count" `
	Price            string  `json:"price" gorm:"type:varchar(10)"`
	Measure          string  `json:"measure" gorm:"type:varchar(10)"`
	PackCount        float32 `json:"packCount"`
	PackUnitQuantity float32 `json:"packUnitQuantity"`
	PackMeasure      string  `json:"packMeasure" gorm:"type:varchar(10)"`
}

func (o OrderItem) Validate() (bool, error) {
	if o.ID == "" {
		return false, errors.New("у одного из продуктов отсутствует id")
	}

	if o.Count <= 0 || o.PackCount <= 0 {
		return false, errors.New(fmt.Sprintf("не корректно указано количество товара, id=%s", o.ID))
	}

	return true, nil
}
