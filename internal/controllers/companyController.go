package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func NewCompany(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user modelsv2.AppUser

		user, ok = u.(modelsv2.AppUser)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		company := modelsv2.Organization{AppUserID: uint(user.ID)}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := json.Unmarshal(data, &company); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		company.ID = 0

		if ok, err := company.Validate(); err != nil || !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := s.DB.Create(&company).Error; err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: company})

	}
}

func UpdateCompany(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user modelsv2.AppUser

		user, ok = u.(modelsv2.AppUser)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		company := modelsv2.Organization{AppUserID: uint(user.ID)}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := json.Unmarshal(data, &company); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if ok, err := company.Validate(); err != nil || !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := s.DB.Select("created_at").First(&company).Error; err != nil {
			log.Println(err.Error())
		}

		if err := s.DB.Save(&company).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: company})
	}
}

func RemoveCompany(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		u, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user modelsv2.AppUser

		user, ok = u.(modelsv2.AppUser)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		company := modelsv2.Organization{AppUserID: uint(user.ID)}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := json.Unmarshal(data, &company); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if ok, err := company.Validate(); err != nil || !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := s.DB.Find(&company).Error; err != nil {
			log.Println(err.Error())
		}

		if company.AppUserID != uint(user.ID) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "permission denied"})
			return
		}

		if err := s.DB.Delete(&company).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: true})
	}
}
