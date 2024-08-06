package apiserver

import (
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
}

func NewServer(dsn string) (*server, error) {
	db, err := configureDatabase(dsn)
	if err != nil {
		return nil, err
	}

	return &server{db: db}, nil
}

func (s *server) Start() {

}
