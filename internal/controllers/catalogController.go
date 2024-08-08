package controllers

import (
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
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
		var products []APiProduct
		res := s.DB.Model(&models.Product{}).Find(&products)
		if res.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: "server error"})
			return
		}

		c.JSON(http.StatusOK, types.Response{Ok: true, Data: products})

	}
}
