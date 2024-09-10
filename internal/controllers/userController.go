package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	models "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

// expect receive initData from Telegram.WebApp.InitData as payload
func GetAppUser(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		ok, err := validateInitData(data, s.Token)
		if !ok || err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		p, _ := url.ParseQuery(string(data))
		us := p.Get("user")

		tgUser := models.TgUser{}

		json.Unmarshal([]byte(us), &tgUser)

		user := models.AppUser{
			TgUser: tgUser,
		}

		if err := s.DB.Preload("Organizations").Preload("Orders").Preload("TgUser").First(&user).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: user})

	}
}

func AppUserUpdate(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		u := models.AppUser{}
		if err := json.Unmarshal(data, &u); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := s.DB.Omit("created_at").Save(&u).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: u})
	}
}
