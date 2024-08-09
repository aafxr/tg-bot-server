package controllers

import (
	"io"
	"net/http"
	"time"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/botserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func PublicPost(s *apiserver.Server, b *botserver.BotServer) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var users []models.TgUser

		if err := s.DB.Find(&users).Error; err != nil && err != gorm.ErrRecordNotFound {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		timer := time.NewTimer(time.Minute * 3)

		go func() {
			<-timer.C

			for _, u := range users {
				msg := tgbotapi.NewMessage(int64(u.ID), string(data))
				b.Bot.Send(msg)
			}
		}()

		ctx.Status(http.StatusOK)
	}
}
