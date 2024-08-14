package modelsv2

type Product struct {
	ID         string     `json:"id" gorm:"column:id;type:varchar(255);not null;primaryKey"`
	ApiCode    string     `json:"apiCode"`
	ApiUID     *string    `json:"apiUID"`
	Currency   string     `json:"currency" gorm:"column:currency;"`
	Price      string     `json:"price" gorm:"column:price;"`
	Preview    string     `json:"preview" gorm:"column:preview;"`
	Photo      []Photo    `json:"photo"`
	Properties []Property `json:"properties"`
}
