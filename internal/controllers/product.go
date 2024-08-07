package controllers

import (
	"net/http"
	"strconv"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func GetProduct(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		prodId := ctx.Param("product_id")
		id, err := strconv.Atoi(prodId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}
		p := models.Product{ID: uint(id)}
		res := s.DB.First(&p)
		if res.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: res.Error.Error()})
			return
		}

		if err := s.DB.Model(&p).Association("Properties").Find(&p.Properties); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: res.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: p})
	}
}
