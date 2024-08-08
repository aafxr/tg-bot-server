package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func GetTGUser(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tgu := models.TgUser{}
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := json.Unmarshal(data, &tgu); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		u := models.AppUser{}
		s.DB.Where("id=?", tgu.AppUserID).First(&u)
		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: u})
	}
}

func GetAppUser(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data := ctx.Param("user_id")
		if data == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unknown user id"})
			return
		}
		id, err := strconv.Atoi(data)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unknown user id"})
			return
		}

		u := models.AppUser{ID: uint(id)}

		if err := s.DB.Model(&u).Omit("Organizations").Preload("TgUser").First(&u).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, u)

	}
}
