package apiserver

import (
	"os"

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
