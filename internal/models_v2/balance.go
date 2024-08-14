package modelsv2

// остатки продукта
type Balance struct {
	ID              uint   `json:"-" gorm:"primaryKey;autoInkrement;not null;unsigned"`
	Quantity        string `gorm:"type:varchar(255)"`
	TradeArea_Id    string `gorm:"type:varchar(255)"`
	TradeArea_Name  string `gorm:"type:varchar(255)"`
	ProductDetailID uint
}
