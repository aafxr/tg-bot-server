package midlewares

import (
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
)

func SessionCheckMW(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var (
			u    interface{}
			user models.AppUser
			ok   bool
		)
		u, ok = ctx.Get("user")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: "need start session"})
			return
		}
		user, ok = u.(models.AppUser)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: "need start session"})
			return
		}
		if user.ID == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: "need start session"})
			return
		}

		ctx.Next()

	}
}
