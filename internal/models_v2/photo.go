package modelsv2

type Photo struct {
	ID        uint   `josn:"id" gorm:"autoIncrement;not null;primaryKey"`
	Src       string `json:"src" gorm:"type:varchar(255)"`
	Preview   bool   `json:"-" gorm:"default:false"`
	ProductID string `json:"-" gorm:"type:varchar(255)"`
}
