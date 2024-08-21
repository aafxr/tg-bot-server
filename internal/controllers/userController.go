package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	models "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// expect receive initData from Telegram.WebApp.InitData as payload
func GetAppUser(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		ok, err := validateInitData(data, s.Token)
		if !ok || err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		p, _ := url.ParseQuery(string(data))
		us := p.Get("user")

		tgUser := models.TgUser{}

		json.Unmarshal([]byte(us), &tgUser)

		user := models.AppUser{
			TgUser: tgUser,
		}

		if err := s.DB.Preload("TgUser").First(&user).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: user})

	}
}

// expect AppUser id named as "uid" in query params
func GetAppUserCompanies(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		uid := ctx.Query("uid")

		id, err := strconv.Atoi(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}
		au := models.AppUser{ID: uint(id)}

		if err := s.DB.Preload("Organizations").First(&au).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: make([]interface{}, 0)})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: au.Organizations})

	}
}

func AppUserNewCompany(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		uid := ctx.Query("uid")
		if uid == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		id, err := strconv.Atoi(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		company := models.Organization{AppUserID: uint(id)}

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

		ok, err := company.Validate()
		if err != nil || !ok {
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

func AppUserUpdateCompany(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		uid := ctx.Query("uid")
		if uid == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		id, err := strconv.Atoi(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		company := models.Organization{AppUserID: uint(id)}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := json.Unmarshal(data, &company); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ok, err := company.Validate()
		if err != nil || !ok {
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

func AppUserRemoveCompany(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		uid := ctx.Query("uid")
		if uid == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		id, err := strconv.Atoi(uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		company := models.Organization{AppUserID: uint(id)}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := json.Unmarshal(data, &company); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ok, err := company.Validate()
		if err != nil || !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		if err := s.DB.Find(&company).Error; err != nil {
			log.Println(err.Error())
		}

		if company.AppUserID != uint(id) {
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
