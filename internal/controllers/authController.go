package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrCorruptedInitData = errors.New("corrupted init data")
)

// receive telegram initData as post payload
/*
ожидает получить telegram initData в теле запроса

после успешной валидации возвращает сгенерированный токен
*/
func AuthTGInitData(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: err.Error()})
			return
		}

		isValid, err := validateInitData(data, s.Token)
		if err != nil || !isValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: err.Error()})
			return
		}

		p, _ := url.ParseQuery(string(data))
		us := p.Get("user")

		u := modelsv2.AppUser{}
		if err := json.Unmarshal([]byte(us), &u.TgUser); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: err.Error()})
			return
		}

		u.TgUser.ID = u.TgUser.ID

		if err := s.DB.Model(&u).Preload("TgUser").Where(&u).First(&u).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, types.Response{Ok: false, Message: err.Error()})
				return
			}
		}

		payload := jwt.MapClaims{
			"uid": u.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		}

		// Создаем новый JWT-токен и подписываем его по алгоритму HS256
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		t, err := token.SignedString([]byte(s.JWTSignKey))
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, types.Response{Ok: false, Message: err.Error()})
			return
		}

		ctx.SetCookie("Authorization", t, 300, "*", s.Domain, false, false)
		ctx.JSON(http.StatusOK, types.Response{Ok: true, Data: t})
	}
}

/*
хелпер, проверяет целостность telegram initdata

returns bool, error
*/
func validateInitData(data []byte, token string) (bool, error) {
	var err error
	var pairs url.Values

	pairs, err = url.ParseQuery(string(data))
	if err != nil {
		return false, err
	}

	hash := pairs.Get("hash")
	pairs.Del("hash")

	if hash == "" {
		return false, ErrCorruptedInitData
	}

	var strs []string
	for k, v := range pairs {
		strs = append(strs, k+"="+v[0])
	}
	sort.Strings(strs)

	authData := strings.Join(strs, "\n")

	enc := sign(authData, token)
	if enc != hash {
		return false, ErrCorruptedInitData

	}

	return true, nil
}

/*
кодирование telegram initData по алгоритму sha256
*/
func sign(authData, token string) string {
	skHmac := hmac.New(sha256.New, []byte("WebAppData"))
	skHmac.Write([]byte(token))

	imrHmac := hmac.New(sha256.New, skHmac.Sum(nil))
	imrHmac.Write([]byte(authData))
	encstr := hex.EncodeToString(imrHmac.Sum(nil))

	return encstr
}
