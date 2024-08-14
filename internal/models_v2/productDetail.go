package modelsv2

type ProductDetail struct {
	ID                       uint   `json:"-" gorm:"primaryKey;autoInkrement;not null;unsigned"`
	ProductID                string `json:"id" gorm:"type:varchar(255);primaryKey;not null"`
	LinkToSite               string `gorm:"type:varchar(255)"`
	PackUnitMeasure          string `gorm:"type:varchar(255)"`
	PackUnitQuantity         string `gorm:"type:varchar(255)"`
	ProductArticle           string `gorm:"type:varchar(255)"`
	ProductArticleForChatBot string `gorm:"type:varchar(255)"`
	ProductId                string `gorm:"type:varchar(255)"`
	ProductName              string `gorm:"type:varchar(255)"`
	TransitAmount            string `gorm:"type:varchar(255)"`
	UnitOfMeasure            string `gorm:"type:varchar(255)"`
	Price_MRC                Price
	Price_RRC                Price
	Balance_Strings          []Balance
	Transit                  Transit
}
