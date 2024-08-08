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
	"github.com/aafxr/tg-bot-server/internal/models"
	"github.com/aafxr/tg-bot-server/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		u := models.AppUser{}
		json.Unmarshal([]byte(us), &u.TgUser)
		u.TgId = u.TgUser.ID

		if err := s.DB.Model(&u).Preload("TgUser").Where(&u).First(&u).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, types.Response{Ok: false, Message: err.Error()})
			return
		}

		session := models.Session{AppUserId: u.ID, TgId: u.TgUser.ID}
		if err := s.DB.Where("tg_id = ?", session.TgId).First(&session).Error; err != nil {
			log.Println(err)
			session.ID = uuid.New().String()
			s.DB.Save(&session)
		}

		ctx.SetCookie(s.SeeeionKey, session.ID, 3600*24*365, "", s.Domain, true, true)
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
