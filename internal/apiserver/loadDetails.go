package apiserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"sync"

	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
)

func (s *Server) LoadDetails() {
	m := make(map[string]bool, 10)
	var prodList []modelsv2.Product
	if err := s.DB.Select("api_code").Find(&prodList).Error; err != nil {
		log.Println(err.Error())
	}
	for _, p := range prodList {
		m[p.ApiCode] = true
	}

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	details := []modelsv2.ProductDetail{}
	sinchronize := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(len(keys))
	for _, k := range keys {
		go loadProducDetails(k, &details, &sinchronize, &wg)
	}

	wg.Wait()

	log.Println("details loaded: ", len(details))

	// for _, d := range details {
	// 	if d.ApiCode == "" {
	// 		continue
	// 	}
	// 	res := s.DB.Save(&d)
	// 	log.Println(d.ProductID, " ", res.Error.Error())

	// }
	// res := s.DB.Save(details)
	// log.Println("details rows affected", res.RowsAffected)

	for _, d := range details {
		if err := s.DB.Create(&d).Error; err != nil {
			log.Println(err.Error())
		}
		// if err := s.DB.Create(&d.Price_MRC).Error; err != nil {
		// 	log.Println(err.Error())
		// }
		// if err := s.DB.Create(&d.Price_RRC).Error; err != nil {
		// 	log.Println(err.Error())
		// }
		// if err := s.DB.Create(&d.Balance_Strings).Error; err != nil {
		// 	log.Println(err.Error())
		// }
		// if err := s.DB.Create(&d.Transit).Error; err != nil {
		// 	log.Println(err.Error())
		// }

	}

	if err := s.DB.Create(&details).Error; err != nil {
		log.Println(err.Error())
	}
}

func loadProducDetails(code string, dt *[]modelsv2.ProductDetail, s *sync.Mutex, wg *sync.WaitGroup) (*[]modelsv2.ProductDetail, error) {
	defer wg.Done()
	if code == "" {
		log.Println(http.ErrContentLength.Error())
		return nil, http.ErrContentLength
	}

	req, err := http.NewRequest("GET", "https://refloor-bot.ru/api/getDetail?", nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	q := req.URL.Query()
	q.Add("code", code)
	req.URL.RawQuery = q.Encode()

	url := req.URL.String()
	details := []modelsv2.ProductDetail{}

	resp, err := http.Get(url)
	log.Println(url, "  ->  ", resp.Status)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	s.Lock()
	defer s.Unlock()

	if data, ok := m["product"]; ok {
		s, err := json.Marshal(data)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		if reflect.ValueOf(data).Kind() == reflect.Slice {
			if err := json.Unmarshal(s, &details); err != nil {
				log.Println(err.Error())
				return nil, err
			}

		} else {
			d := modelsv2.ProductDetail{}
			if err := json.Unmarshal(s, &d); err != nil {
				log.Println(err.Error())
				return nil, err
			}
			details = append(details, d)
		}

	} else {
		log.Println("product not found")
	}

	for _, d := range details {
		d.ApiCode = code
	}

	*dt = append(*dt, details...)

	return &details, nil
}
