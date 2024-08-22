package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
	"github.com/aafxr/tg-bot-server/internal/botserver"
	"github.com/aafxr/tg-bot-server/internal/controllers"
	"github.com/aafxr/tg-bot-server/internal/midlewares"
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
	// gin.DisableConsoleColor()

	// // Logging to a file.
	// f, _ := os.Create("gin.log")
	// defer f.Close()
	// gin.DefaultWriter = io.MultiWriter(f)

	s, err := apiserver.NewServer(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// gin.SetMode(gin.ReleaseMode)
	s.LoadDBData()

	b, err := botserver.NewBotServer(s)
	if err != nil {
		log.Fatal(err)
	}

	go b.Run()

	baseRouter := gin.New()

	baseRouter.Use(midlewares.CORSMiddleware())
	baseRouter.Use(gin.Logger())
	baseRouter.Use(gin.Recovery())

	r := baseRouter.Group("/api")

	r.Use(midlewares.UserLoadMW(s))

	r.Static("/assets", "./assets")

	r.GET("/catalog", controllers.GetCatalogHandler(s))
	r.GET("/product/:product_id", controllers.GetProduct(s))
	r.GET("/products", controllers.GetProductsList(s))

	r.POST("/me", controllers.GetAppUser(s))
	r.POST("/user/update", controllers.AppUserUpdate(s))

	r.POST("/order/new", controllers.NewOrder(s))

	r.GET("/companies", controllers.GetAppUserCompanies(s))
	r.POST("/company/new", controllers.AppUserNewCompany(s))
	r.POST("/company/update", controllers.AppUserUpdateCompany(s))
	r.POST("/company/remove", controllers.AppUserRemoveCompany(s))

	if err := baseRouter.Run(os.Getenv("DOMAIN")); err != nil {
		log.Fatal(err)
	}

}
