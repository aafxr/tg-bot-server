package controllers

import (
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

type APiProduct struct {
	ID       uint
	Title    string
	Currency string
	Price    uint
	Preview  string
}

func GetCatalogHandler(s *apiserver.Server) func(*gin.Context) {
	return func(c *gin.Context) {
		result := make(map[string]interface{})

		var products []modelsv2.Product
		if err := s.DB.Preload("Properties").Preload("Photo").Find(&products).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		result["elements"] = products

		articles := []modelsv2.Article{}
		if err := s.DB.Find(&articles).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		am := make(map[string]string, len(articles))
		for _, el := range articles {
			am[el.Name] = el.ProductID
		}

		result["article"] = am

		sections := []modelsv2.Section{}
		if err := s.DB.Find(&sections).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		result["sections"] = sections

		c.JSON(http.StatusOK, types.Response{Ok: true, Data: result})

	}
}
