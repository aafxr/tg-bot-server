package modelsv2

type Property struct {
	ID        uint   `json:"-" gorm:"autoIncrement;primaryKey;unsigned"`
	ProductID string `json:"id" gorm:"type:varchar(255)"`
	Name      string `json:"name" gorm:"name;type:varchar(255)"`
	Value     string `json:"value" gorm:"value;type:varchar(255)"`
	Code      string `json:"code" gorm:"type:varchar(255)"`
}
