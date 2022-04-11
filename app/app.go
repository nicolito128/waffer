package app

import (
	"log"

	"github.com/nicolito128/waffer/pkg/bot"
)

var Bot = StartSession()

func Init() {
	Bot.Run()
}

func StartSession() *bot.Bot {
	bot, err := bot.New()
	if err != nil {
		log.Fatalf("error cannot create a new app-bot: %s", err)
	}

	return bot
}
