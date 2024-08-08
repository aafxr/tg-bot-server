package models

type Session struct {
	ID        string
	TgId      uint
	AppUserId uint
	AppUser   AppUser
}
