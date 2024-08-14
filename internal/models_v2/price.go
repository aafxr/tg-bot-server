package modelsv2

//mrc / rrc
type Price struct {
	ID              uint   `json:"-" gorm:"primaryKey;autoInkrement;not null;unsigned"`
	Name            string `gorm:"type:varchar(255)"`
	UnitOfMeasure   string `gorm:"type:varchar(255)"`
	Value           string `gorm:"type:varchar(255)"`
	ProductDetailID uint
}
