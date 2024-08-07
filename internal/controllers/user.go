package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func GetUser(s *apiserver.Server) func(*gin.Context) {
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
