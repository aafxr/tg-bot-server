package botserver

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type BotServer struct {
	Token string
	Bot   *tgbotapi.BotAPI
	s     *apiserver.Server
}

func NewBotServer(s *apiserver.Server) (*BotServer, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, errors.New("token not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bs := &BotServer{Token: token, Bot: bot, s: s}
	return bs, nil
}

func (b *BotServer) Run() {

	// b.Bot.Debug = true

	log.Printf("Authorized on account %s", b.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				b.handleStart(update)
			case "help":
				b.handleHelp(update)
			}
			continue
		}

		b.handleTextMessage(update)
	}
}

func (b *BotServer) handleStart(update tgbotapi.Update) error {
	u := update.SentFrom()
	appUser := models.AppUser{TgId: uint(u.ID)}

	res := b.s.DB.First(&appUser)
	if res.Error != nil {
		if res.Error != gorm.ErrRecordNotFound {
			return res.Error
		}
		tgu := models.TgUser{
			ID:        uint(u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.UserName,
		}
		appUser.TgUser = tgu
		res := b.s.DB.Create(&appUser)
		if res.Error != nil {
			return res.Error
		}
	}

	if err := b.s.DB.Model(&appUser).Association("TgUser").Find(&appUser.TgUser); err != nil {
		return err
	}

	text := fmt.Sprintf("hello %s %s %s", appUser.TgUser.FirstName, appUser.TgUser.LastName, appUser.TgUser.Nickname)
	msg := tgbotapi.NewMessage(update.FromChat().ID, text)
	b.Bot.Send(msg)

	return nil
}

func (b *BotServer) handleHelp(update tgbotapi.Update) {
	text := `
	доступны команды:
	/start
	/help
	`

	chat := update.FromChat()
	msg := tgbotapi.NewMessage(chat.ID, text)
	b.Bot.Send(msg)
}

func (b *BotServer) handleTextMessage(update tgbotapi.Update) {
	logText("handleTextMessage " + update.Message.Text)

}

func logText(t string) {
	log.Println(t)
}
