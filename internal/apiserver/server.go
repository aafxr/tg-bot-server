package apiserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"gorm.io/gorm"
)

type Server struct {
	DB         *gorm.DB
	Token      string
	SeeeionKey string
	Domain     string
}

func NewServer(dsn string) (*Server, error) {
	db, err := configureDatabase(dsn)
	if err != nil {
		return nil, err
	}

	return &Server{
		DB:         db,
		Token:      os.Getenv("BOT_TOKEN"),
		SeeeionKey: os.Getenv("SESSEION_KEY"),
		Domain:     os.Getenv("DOMAIN"),
	}, nil
}

func (s *Server) Start() {

}

func (s *Server) LoadDBData() {
	LoadProducts(s)
}

// load products from [https://refloor-opt.ru/api/telegram/]
func LoadProducts(s *Server) {
	resp, e := http.Get("https://refloor-opt.ru/api/telegram/")
	if e == nil {
		body, e := io.ReadAll(resp.Body)
		if e == nil {
			var result map[string]interface{}
			if e := json.Unmarshal(body, &result); e == nil {

				if elMap, ok := result["elements"].(map[string]interface{}); ok {
					elements := make([]modelsv2.Product, len(elMap))
					var photoID uint = 1
					for k, v := range elMap {
						var item modelsv2.Product
						d, e := json.Marshal(v)
						if e != nil {
							log.Println(k, " - failed")
							continue
						}
						if e := json.Unmarshal(d, &item); e != nil {
							log.Println(k, " - failed")
							continue
						}

						for _, p := range item.Photo {
							p.ProductID = item.ID
							p.ID = photoID
							photoID++
							src, err := loadPhoto(p.Src)
							if err == nil {
								p.Src = fmt.Sprintf("%s/api/assets/%s", s.Domain, src)
							} else {
								log.Println(err.Error())
							}
						}

						src, err := loadPhoto(item.Preview)
						if err == nil {
							item.Preview = fmt.Sprintf("%s/api/assets/%s", s.Domain, src)
						} else {
							log.Println(err.Error())
						}

						idx := slices.IndexFunc(item.Photo, func(el modelsv2.Photo) bool { return el.Src == item.Preview })
						if idx != -1 {
							item.Photo[idx].Preview = true
						}

						elements = append(elements, item)
					}

					for _, el := range elements {
						if err := s.DB.Save(&el).Error; err != nil {
							log.Println(err.Error())
						}
					}
				} else {
					log.Println("casting failed===================")
				}
			} else {
				log.Println("[https://refloor-opt.ru/api/telegram/] ", e.Error())
			}

		} else {
			log.Println("[https://refloor-opt.ru/api/telegram/] ", e.Error())
		}
	} else {
		log.Println("[https://refloor-opt.ru/api/telegram/] ", e.Error())
	}
}

func loadPhoto(s string) (string, error) {
	sl := strings.Split(s, "/")
	fileName := sl[len(sl)-1]

	if fileName == "" {
		return "", errors.New("file name empty")
	}

	relativePath := "assets/" + fileName

	dir := filepath.Dir(relativePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(relativePath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(relativePath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		resp, err := http.Get(s)
		if err != nil {
			return "", err
		}

		_, e := io.Copy(file, resp.Body)
		if e != nil {
			return "", e
		}
	}

	return fileName, nil
}
