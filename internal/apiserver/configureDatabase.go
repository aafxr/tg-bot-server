package apiserver

import (
	"time"

	"github.com/aafxr/tg-bot-server/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func configureDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(
		&models.AppUser{},
		&models.TgUser{},
		&models.Organization{},
		&models.Product{},
		&models.ProductProperty{},
		&models.Order{},
		&models.OrderItem{},
	)

	return db, nil
}
