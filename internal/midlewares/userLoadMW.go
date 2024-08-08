package midlewares

import (
	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/gin-gonic/gin"
)

// предварительная загрузка пользователя в контекст
func UserLoadMW(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		session := models.Session{}

		sesId, err := ctx.Cookie(s.SeeeionKey)
		if err != nil {
			ctx.Next()
			return
		}

		session.ID = sesId
		if err := s.DB.Model(&session).Preload("AppUser").Preload("AppUser.TgUser").Find(&session).Error; err != nil {
			ctx.Next()
			return
		}

		ctx.Set("user", session.AppUser)
		ctx.Next()
	}
}
