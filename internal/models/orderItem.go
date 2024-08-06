package models

type OrderItem struct {
	ID        uint   `json:"id" gorm:"id;type:int;unsigned;auto increment;primaryKey;"`
	Count     uint   `json:"count" gorm:"count"`
	Price     uint   `json:"price" gorm:"price"`
	Currency  string `json:"currency" gorm:"currency"`
	OrderID   uint
	ProductID uint
}
