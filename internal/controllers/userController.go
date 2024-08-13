package controllers

import (
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	models "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func GetAppUser(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}
		var user models.AppUser
		user, ok = data.(models.AppUser)

		ctx.JSON(http.StatusOK, user)

	}
}
