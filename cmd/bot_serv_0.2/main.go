package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/botserver"
	"github.com/aafxr/tg-bot-server/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var dsn string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("DB_USER_NAME")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, pass, host, port, dbName)

}

func main() {
	s, err := apiserver.NewServer(dsn)
	if err != nil {
		log.Fatal(err)
	}

	b, err := botserver.NewBotServer(s)
	if err != nil {
		log.Fatal(err)
	}

	go b.Run()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/catalog", controllers.GetCatalogHandler(s))
	r.GET("/catalog/:product_id/details", controllers.GetProduct(s))

	//info about user
	r.POST("/user", controllers.GetUser(s))

	if err := r.Run(os.Getenv("DOMAIN")); err != nil {
		log.Fatal(err)
	}

}
