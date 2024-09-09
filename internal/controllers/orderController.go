package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

// expect uid in query params
/*
method - post
returns - возврацает созданный заказ с обновленным полем id
*/
func NewOrder(s *apiserver.Server) func(*gin.Context) {
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

		u := modelsv2.AppUser{ID: uint(id)}
		if err := s.DB.First(&u).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: "unauthorizet"})
			return
		}

		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		o := modelsv2.Order{}
		if err := json.Unmarshal(data, &o); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
			return
		}

		_, e := o.Validate()
		if e != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: e.Error()})
			return
		}

		if err := s.DB.Save(&o).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: o})
	}
}
