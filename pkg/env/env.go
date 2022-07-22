package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadENV() {
	if os.Getenv("BOT_MODE") == "production" {
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get(key string) string {
	loadENV()
	return os.Getenv(key)
}
