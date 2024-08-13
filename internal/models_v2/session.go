package modelsv2

type Session struct {
	ID        string
	TgUserID  uint
	AppUserID uint
	AppUser   AppUser
}
