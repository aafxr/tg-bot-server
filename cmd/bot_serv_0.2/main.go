package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aafxr/tg-bot-server/internal/apiserver"
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
	s.Start()

}
