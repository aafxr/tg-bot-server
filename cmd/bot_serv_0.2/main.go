package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/botserver"
	"github.com/aafxr/tg-bot-server/internal/controllers"
	"github.com/aafxr/tg-bot-server/internal/midlewares"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			slices.ContainsFunc([]string{"localhost", "postman"}, func(s string) bool {
				return strings.Contains(origin, s)
			})
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(midlewares.UserLoadMW(s))

	r.Static("/assets", "./assets")

	r.GET("/catalog", controllers.GetCatalogHandler(s))
	r.GET("/catalog/:product_id/details", controllers.GetProduct(s))

	r.POST("/session", controllers.StartSession(s))

	r.GET("/test", controllers.Test)

	authRouter := r.Group("")
	authRouter.Use(midlewares.SessionCheckMW(s))
	{
		// authRouter.POST("/user", controllers.GetTGUser(s))
		authRouter.GET("/me", controllers.GetAppUser(s))
		authRouter.GET("/myOrganizations", controllers.GetUserOrganizations(s))
		authRouter.POST("/publishPost", controllers.PublicPost(s, b))

	}

	if err := r.Run(os.Getenv("DOMAIN")); err != nil {
		log.Fatal(err)
	}
}
