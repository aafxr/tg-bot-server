package modelsv2

type Session struct {
	ID        string
	TgId      uint
	AppUserID uint
	AppUser   AppUser
}
