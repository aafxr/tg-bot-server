package modelsv2

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	ID        string `json:"id" gorm:"id;type:varchar(255);primaryKey;"`
	Name      string `json:"name" gorm:"name"`
	Value     string `json:"value" gorm:"value"`
	ProductID string `gorm:"type:varchar(255)"`
}
