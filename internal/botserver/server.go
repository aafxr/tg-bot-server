package botserver

import (
	"errors"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotServer struct {
	Token string
	Bot   *tgbotapi.BotAPI
}

func NewBotServer() (*BotServer, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return nil, errors.New("token not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		return nil, err
	}
	bs := &BotServer{Token: token, Bot: bot}
	return bs, nil
}

func (b *BotServer) Run() {

	b.Bot.Debug = true

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

func (b *BotServer) handleStart(u tgbotapi.Update) {

}

func (b *BotServer) handleHelp(u tgbotapi.Update) {

}

func (b *BotServer) handleTextMessage(u tgbotapi.Update) {

}
