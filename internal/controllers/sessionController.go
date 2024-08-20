package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrCorruptedInitData = errors.New("corrupted init data")
)

func StartSession(s *apiserver.Server) func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, types.Response{Ok: false, Message: err.Error()})
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
		json.Unmarshal([]byte(us), &u.TgUser)
		u.TgUser.ID = u.TgUser.ID

		if err := s.DB.Model(&u).Preload("TgUser").Where(&u).First(&u).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				// ctx.AbortWithStatusJSON(http.StatusNotFound, types.Response{Ok: false, Message: "User not found. Type /start in telegram chat bot "})
				// return
				ctx.AbortWithStatusJSON(http.StatusNotFound, types.Response{Ok: false, Message: err.Error()})
				return
			}
		}

		session := modelsv2.Session{AppUserID: u.ID, TgUserID: u.TgUser.ID}
		if err := s.DB.Where("tg_user_id = ?", session.TgUserID).First(&session).Error; err != nil {
			log.Println(err)
			session.ID = uuid.New().String()
			s.DB.Omit(clause.Associations).Create(&session)
		}

		ctx.SetCookie(s.SeeeionKey, session.ID, 3600*24*365, "", s.Domain, false, false)
		ctx.JSON(http.StatusOK, types.Response{Ok: true})
	}
}

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

func sign(authData, token string) string {
	skHmac := hmac.New(sha256.New, []byte("WebAppData"))
	skHmac.Write([]byte(token))

	imrHmac := hmac.New(sha256.New, skHmac.Sum(nil))
	imrHmac.Write([]byte(authData))
	encstr := hex.EncodeToString(imrHmac.Sum(nil))

	return encstr
}
