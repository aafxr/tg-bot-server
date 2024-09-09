package apiserver

import (
	"os"

	"gorm.io/gorm"
)

// # Domain - домен
//
// # Token - токен бота
//
// JWTSignKey - используется для генерации токена пользователя
type Server struct {
	DB         *gorm.DB
	Token      string
	SeeeionKey string
	Domain     string
	JWTSignKey string
}

// метод конфигурирует бд и возвращает инстанс суцности Server
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
		JWTSignKey: "zhfunssnqfjbpidtrokvsksrspxfhzxelxsruegjawfkyhqgyu",
	}, nil
}

func (s *Server) Start() {

}

/*
метод для загрузки каталога продуктов и деталей о товаре
*/
func (s *Server) LoadDBData() {
	// s.LoadProducts()
	// s.LoadDetails()

}
