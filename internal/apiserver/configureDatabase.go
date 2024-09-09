package apiserver

import (
	"time"

	modelsv2 "github.com/aafxr/tg-bot-server/internal/models_v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// устанавливат соединение с бд,
//
// настраивает параметры подключения к бд
//
// выполняет синхронизацию сущностей приложения и таблиц в бд
//
// dsn - строка подключения к бд вида: user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True
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
		&modelsv2.AppUser{},
		&modelsv2.TgUser{},
		&modelsv2.Order{},
		&modelsv2.OrderItem{},
		&modelsv2.Organization{},
		//-----------------------
		&modelsv2.Session{},
		//-----------------------
		&modelsv2.Product{},
		// &modelsv2.Property{},
		// &modelsv2.Photo{},
		&modelsv2.Article{},
		&modelsv2.Section{},
		&modelsv2.Storehouse{},
		&modelsv2.StorehouseProduct{},
		//-----------------------
		&modelsv2.ProductDetail{},
		// &modelsv2.Balance{},
		// &modelsv2.Price{},
		// &modelsv2.Transit{},
	)

	return db, nil
}
