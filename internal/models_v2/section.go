package modelsv2

type Section struct {
	ID     uint     `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Items  []string `json:"items" gorm:"type:json;serializer:json"`
	Parent string   `json:"parent" gorm:"type:varchar(255)"`
	Sort   string   `json:"sort" gorm:"type:varchar(255)"`
	Title  string   `json:"title" gorm:"type:varchar(255)"`
}
