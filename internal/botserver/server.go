package botserver

import (
	"errors"
	"log"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotServer struct {
	Token string
	Bot   *tgbotapi.BotAPI
	s     *apiserver.Server
}

func NewBotServer(s *apiserver.Server) (*BotServer, error) {
	token := s.Token
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
				botServLog("./start", b.handleStart(update))
			case "help":
				botServLog("/help", b.handleHelp(update))
			case "companies":
				botServLog("/companies", b.handleCompanies(update))
			default:
				botServLog("other", b.handleOtherComands(update))
			}
			continue
		}

		botServLog(update.Message.Text, b.handleTextMessage(update))
	}
}
