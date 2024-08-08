package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	v, e := ctx.Get("user")
	if e {
		ctx.JSON(http.StatusOK, v)
	}

}
