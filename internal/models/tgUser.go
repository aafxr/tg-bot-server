package models

type TgUser struct {
	ID        uint   `json:"id"         gorm:"column:id;type:int;unsigned;auto increment;not null;primaryKey"`
	FirstName string `json:"first_name" gorm:"column:first_name;type:varchar(255)"`
	LastName  string `json:"last_name"  gorm:"column:last_name;type:varchar(255)"`
	Nickname  string `json:"nickname"   gorm:"column:nickname;type:varchar(255)"`
	Photo     string `json:"photo"      gorm:"column:photo;type:varchar(255)"`
	AppUserID uint
}
