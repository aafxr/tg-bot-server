package controllers

import (
	"net/http"
	"slices"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	models "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func GetProduct(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		prodId := ctx.Param("product_id")
		if prodId == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "missing product_id in uri"})
			return
		}
		p := models.Product{ID: prodId}
		res := s.DB.Preload("Photo").Preload("Properties").First(&p)
		if res.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: res.Error.Error()})
			return
		}

		// if err := s.DB.Model(&p).Preload("Photos").Preload("Properties").Find(&p); err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: res.Error.Error()})
		// 	return
		// }

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: p})
	}
}

func GetProductsList(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var products []models.Product
		if e := s.DB.Preload("Photo").Preload("Properties").Find(&products).Error; e != nil {
			ctx.JSON(http.StatusInternalServerError, types.Response{Ok: false, Message: e.Error()})
			return
		}

		for _, p := range products {
			idx := slices.IndexFunc(p.Photo, func(el models.Photo) bool { return el.Preview })
			if idx != -1 {
				p.Preview = p.Photo[idx].Src
			}
		}
		ctx.JSON(http.StatusOK, products)
	}
}
