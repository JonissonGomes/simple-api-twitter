package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

type Handler struct {
	DB *mgo.Session
}

var DbKey = getKey("PASSWORD")

func getKey(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Erro ao carregar .env file")
	}

	return os.Getenv(key)
}
