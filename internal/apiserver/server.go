package apiserver

import (
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func NewServer(dsn string) (*Server, error) {
	db, err := configureDatabase(dsn)
	if err != nil {
		return nil, err
	}

	return &Server{DB: db}, nil
}

func (s *Server) Start() {

}
