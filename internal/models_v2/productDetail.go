package modelsv2

type ProductDetail struct {
	ID                       uint      `json:"-" gorm:"primaryKey;autoInkrement;not null;unsigned"`
	ApiCode                  string    `json:"-" gorm:"type:varchar(255)"`
	ProductID                string    `json:"ProductId" gorm:"type:varchar(255)"`
	LinkToSite               string    `json:"LinkToSite" gorm:"type:varchar(255)"`
	PackUnitMeasure          string    `json:"PackUnitMeasure" gorm:"type:varchar(255)"`
	PackUnitQuantity         string    `json:"PackUnitQuantity" gorm:"type:varchar(255)"`
	ProductArticle           string    `json:"ProductArticle" gorm:"type:varchar(255)"`
	ProductArticleForChatBot string    `json:"ProductArticleForChatBot" gorm:"type:varchar(255)"`
	ProductName              string    `json:"ProductName" gorm:"type:varchar(255)"`
	TransitAmount            string    `json:"TransitAmount" gorm:"type:varchar(255)"`
	UnitOfMeasure            string    `json:"UnitOfMeasure" gorm:"type:varchar(255)"`
	Price_MRC                Price     `json:"Price_MRC" gorm:""`
	Price_RRC                Price     `json:"Price_RRC" gorm:""`
	Balance_Strings          []Balance `json:"Balance_Strings" gorm:""`
	Transit                  Transit   `json:"Transit" gorm:""`
}
