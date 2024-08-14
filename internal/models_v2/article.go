package modelsv2

type Article struct {
	ID        uint   `json:"-" gorm:"primaryKey;autoIncrement;not null;"`
	ProductID string `gorm:"type:varchar(255)"`
	Name      string `json:"article" gorm:"type:varchar(255)"`
}
