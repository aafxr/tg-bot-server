package main

import (
	"fmt"
	"log"
	"os"
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
	// gin.DisableConsoleColor()

	// // Logging to a file.
	// f, _ := os.Create("gin.log")
	// defer f.Close()
	// gin.DefaultWriter = io.MultiWriter(f)

	s, err := apiserver.NewServer(dsn)
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	s.LoadDBData()

	b, err := botserver.NewBotServer(s)
	if err != nil {
		log.Fatal(err)
	}

	go b.Run()

	baseRouter := gin.New()
	r := baseRouter.Group("/api")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return slices.ContainsFunc([]string{"localhost", "postman", "127.0.0.1"}, func(s string) bool {
		// 		return strings.Contains(origin, s)
		// 	})
		// },
		MaxAge: 12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(midlewares.UserLoadMW(s))

	r.Static("/assets", "./assets")

	// r.GET("/catalog", controllers.GetCatalogHandler(s))
	r.GET("/product/:product_id", controllers.GetProduct(s))
	r.GET("/products", controllers.GetProductsList(s))

	// r.POST("/session", controllers.StartSession(s))

	// r.GET("/test", controllers.Test)
	// r.POST("/upload", controllers.UploadFile(s))

	// authRouter := r.Group("")
	// authRouter.Use(midlewares.SessionCheckMW(s))
	// {
	// 	// authRouter.POST("/user", controllers.GetTGUser(s))
	// 	authRouter.GET("/me", controllers.GetAppUser(s))
	// 	authRouter.POST("/newOrganization", controllers.NewOrganization(s))
	// 	authRouter.GET("/myOrganizations", controllers.GetUserOrganizations(s))
	// 	authRouter.POST("/publishPost", controllers.PublicPost(s, b))

	// }

	if err := baseRouter.Run(os.Getenv("DOMAIN")); err != nil {
		log.Fatal(err)
	}

}
