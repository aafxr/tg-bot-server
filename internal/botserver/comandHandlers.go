package botserver

import (
	"fmt"
	"log"

	models "github.com/aafxr/tg-bot-server/internal/models_v2"
	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func (b *BotServer) handleStart(update tgbotapi.Update) error {
	u := update.SentFrom()
	appUser := models.AppUser{TgUser: modelsv2.TgUser{ID: uint(u.ID)}}

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

func (b *BotServer) handleHelp(update tgbotapi.Update) error {
	text := `
	доступны команды:
	/start
	/help
	`

	chat := update.FromChat()
	msg := tgbotapi.NewMessage(chat.ID, text)
	_, err := b.Bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *BotServer) handleCompanies(update tgbotapi.Update) error {
	chatId := update.FromChat().ID
	var msg tgbotapi.MessageConfig
	u := modelsv2.AppUser{TgUser: modelsv2.TgUser{ID: uint(update.SentFrom().ID)}}

	if err := b.s.DB.Model(&u).Preload("Organizations").Where(&u).First(&u).Error; err != nil {
		msg = tgbotapi.NewMessage(chatId, "Не удалось найти запись о пользоватле")
		if _, e := b.Bot.Send(msg); e != nil {
			return e
		}
		return nil
	}

	if len(u.Organizations) == 0 {
		msg = tgbotapi.NewMessage(chatId, "Список компаний пуст")
		if _, e := b.Bot.Send(msg); e != nil {
			return e
		}
		return nil
	}

	text := "Ваши компании:\n"
	for _, o := range u.Organizations {
		text += fmt.Sprintf("Компания: %s \nГород: %s\nСтрана: %s\n\n", o.Name, o.City, o.Country)
	}

	msg = tgbotapi.NewMessage(chatId, text)
	if _, e := b.Bot.Send(msg); e != nil {
		return e
	}

	return nil
}

func (b *BotServer) handleOtherComands(update tgbotapi.Update) error {
	log.Println("[bot server] command:", update.Message.Text)
	return nil
}
