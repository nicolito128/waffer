package app

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/bot"
	"github.com/nicolito128/waffer/pkg/config"
	"github.com/nicolito128/waffer/pkg/env"
	"github.com/nicolito128/waffer/stdcommands"
)

var intents = discordgo.IntentsGuilds |
	discordgo.IntentsGuildMessages |
	discordgo.IntentDirectMessageTyping |
	discordgo.IntentDirectMessages |
	discordgo.IntentGuildMessages |
	discordgo.IntentsMessageContent |
	discordgo.IntentGuildMembers |
	discordgo.IntentsAllWithoutPrivileged |
	discordgo.IntentGuildPresences |
	discordgo.IntentsAllWithoutPrivileged |
	discordgo.IntentsGuilds

func Start() {
	s, err := bot.New(&config.ConnectionConfig{
		Token:   env.Get("BOT_TOKEN"),
		Intents: intents,
		OwnerID: env.Get("OWNER_ID"),
		Prefix:  env.Get("BOT_PREFIX"),
	})

	if err != nil {
		log.Fatalf("Error creating Discord session: %s", err)
	}

	s.Identify.Presence.Game.Name = env.Get("BOT_PREFIX")

	stdcommands.LoadCommands(s)

	if err = bot.Init(s); err != nil {
		log.Fatalf("Error initializing bot: %s", err)
	}
}
