package controllers

import (
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	models "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

/*
method - get

returns - список созанных пользователем организаций
*/
func GetUserOrganizations(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user models.AppUser

		user, ok = u.(models.AppUser)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var orgs []models.Organization
		if err := s.DB.Where("app_user_id = ?", user.ID).Find(&orgs).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: orgs})
	}
}

func NewOrganization(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, types.Response{Ok: true})
	}
}
