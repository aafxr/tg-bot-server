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
	"reflect"
	"slices"
	"strings"

	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
)

// load products from [https://refloor-opt.ru/api/telegram/]
// сохраняет полученные продукты в локальную бд
func (s *Server) LoadProducts() {
	resp, e := http.Get("https://refloor-opt.ru/api/telegram/")
	if e == nil {
		body, e := io.ReadAll(resp.Body)
		if e == nil {
			var result map[string]interface{}
			if e := json.Unmarshal(body, &result); e == nil {

				if elMap, ok := result["elements"].(map[string]interface{}); ok {
					elements := make([]modelsv2.Product, len(elMap))
					var photoID uint = 1
					var propID uint = 1
					var item modelsv2.Product
					for k, v := range elMap {
						item = modelsv2.Product{}
						d, e := json.Marshal(v)
						if e != nil {
							log.Println(k, " - failed")
							continue
						}
						if e := json.Unmarshal(d, &item); e != nil {
							log.Println(k, " - failed")
							continue
						}

						for i, _ := range item.Photo {
							item.Photo[i].ProductID = item.ID
							item.Photo[i].ID = photoID
							photoID++
							src, err := loadPhoto(item.Photo[i].Src)
							if err == nil {
								item.Photo[i].Src = fmt.Sprintf("%s/api/assets/%s", s.Domain, src)
							} else {
								log.Println(err.Error())
							}
						}

						for i := range item.Properties {
							item.Properties[i].ID = propID
							propID++
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

						if item.ID != "" {
							elements = append(elements, item)
						}
					}

					// for _, e := range elements {
					// 	if e.ID == "" {
					// 		continue
					// 	}
					// 	res := s.DB.Save(&e)
					// 	if res.Error != nil {
					// 		log.Println(e.ID, " ", res.Error.Error())
					// 	}

					// }

					// res := s.DB.Save(elements)

					// log.Println("RowsAffected ", res.RowsAffected)

					for _, el := range elements {
						if err := s.DB.Create(&el).Error; err != nil {
							log.Println(err.Error())
							// if len(el.Properties) > 0 {
							// 	if err := s.DB.Create(&el.Properties).Error; err != nil {
							// 		log.Println(err.Error())
							// 	}
							// }
							// if len(el.Photo) > 0 {
							// 	if err := s.DB.Create(&el.Photo).Error; err != nil {
							// 		log.Println(err.Error())
							// 	}
							// }
						}
					}
				} else {
					log.Println("casting failed===================")
				}

				/*

					---- articles ----

				*/
				if arMap, ok := result["articles"].(map[string]interface{}); ok {
					var arIdx uint = 1
					for k, v := range arMap {
						val := ""
						d, e := json.Marshal(v)
						if e != nil {
							log.Println(e.Error())
							continue
						}

						if e := json.Unmarshal(d, &val); e != nil {
							log.Println(e.Error())
							continue
						}
						a := modelsv2.Article{ID: arIdx, ProductID: val, Name: k}
						arIdx++
						if e := s.DB.Create(&a).Error; e != nil {
							log.Println(e.Error())
						}
					}
				} else {
					log.Println("casting failed=================== articles")
				}

				/*

					---- sections ----

				*/
				log.Println(result["sections"])
				if sMap, ok := result["sections"].([]interface{}); ok {
					for _, v := range sMap {
						sec := modelsv2.Section{}

						parent := v.(map[string]interface{})["parent"]
						if reflect.ValueOf(parent).Kind() != reflect.String {
							v.(map[string]interface{})["parent"] = ""
						}
						d, e := json.Marshal(v)
						if e != nil {
							log.Println(e.Error())
							continue
						}

						if e := json.Unmarshal(d, &sec); e != nil {
							log.Println(e.Error())
							continue
						}

						if e := s.DB.Create(&sec).Error; e != nil {
							log.Println(e.Error())
						}
					}
				} else {
					log.Println("casting failed=================== sections")
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

/*
скачивание изображений в локальную бд
*/
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
