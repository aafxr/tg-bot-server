package modelsv2

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"column:name;type:varchar(255)"`
	FullName  string `json:"fullName" gorm:"column:full_name;type:varchar(255)"`
	Address   string `json:"address" gorm:"column:address;type:varchar(255)"`
	Country   string `json:"country" gorm:"column:country;type:varchar(255)"`
	City      string `json:"city" gorm:"column:city;type:varchar(255)"`
	INN       string `json:"INN" gorm:"column:inn;type:varchar(10)"`
	AppUserID uint
}

func (o Organization) Validate() (bool, error) {
	matched, _ := regexp.MatchString(`^\d$`, "Blackcat meow")
	if !matched && len(o.INN) != 10 {
		return false, errors.New("не корректно указан ИНН")
	}

	if o.Address == "" {
		return false, errors.New("необходимо указать адресс")
	}

	if o.Name == "" {
		return false, errors.New("необходимо указать имя организации")
	}

	if o.FullName == "" {
		return false, errors.New("Необходимо указать полное имя организации")
	}

	return true, nil
}
