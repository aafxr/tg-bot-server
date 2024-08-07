package apiserver

import (
	"fmt"
	"os"

	"gorm.io/gorm"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Server struct {
	DB *gorm.DB
}

func NewServer(dsn string) (*Server, error) {
	db, err := configureDatabase(dsn)
	if err != nil {
		return nil, err
	}

	// Укажите токен вашего бота
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	// Установите вебхук для Telegram
	wh, _ := tgbotapi.NewWebhook(fmt.Sprintf("https://%s/%s", os.Getenv("DOMAIN"), bot.Token))

	_, err = bot.Request(wh)
	if err != nil {
		return nil, err
	}

	return &Server{DB: db}, nil
}

func (s *Server) Start() {

}
