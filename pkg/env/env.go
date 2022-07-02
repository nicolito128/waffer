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

func GetBotMode() string {
	loadENV()
	return os.Getenv("BOT_MODE")
}

func GetOwnerID() string {
	loadENV()
	return os.Getenv("OWNER_ID")
}

func GetBotToken() string {
	loadENV()
	return os.Getenv("BOT_TOKEN")
}

func GetBotPrefix() string {
	loadENV()
	return os.Getenv("BOT_PREFIX")
}
