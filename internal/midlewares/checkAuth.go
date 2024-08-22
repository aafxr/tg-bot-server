package midlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckAuth(s *apiserver.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString, err := ctx.Cookie("Authorization")
		if err != nil {
			authHeader := ctx.Request.Header["Authorization"]
			if len(authHeader) > 0 {
				tokenString = authHeader[0]
			}
			if tokenString == "" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: err.Error()})
				return
			}
		}

		t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(s.JWTSignKey), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: err.Error()})
			return
		}

		uid := t.Header["uid"]

		ctx.Set("uid", uid)

		log.Println("uid: ", uid)

	}
}
