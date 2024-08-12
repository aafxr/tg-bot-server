package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type R struct {
	Name   string
	Val    string
	Client map[string]interface{}
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": R{Name: "test", Val: "123", Client: map[string]interface{}{"a": 12, "b": "wd"}},
	})

}
