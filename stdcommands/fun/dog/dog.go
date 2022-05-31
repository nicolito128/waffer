package dog

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/queries"
	"github.com/nicolito128/waffer/plugins/commands"
)

type DogAPI struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var api = "https://dog.ceo/api/breeds/image/random"

var Command = &commands.WafferCommand{
	Name:        "dog",
	Description: "Sends a random dog image",
	Aliases:     []string{"doggo"},
	Category:    "fun",

	RunInDM: true,

	Arguments:    []string{},
	RequiredArgs: 0,

	DiscordPermissions: discordgo.PermissionSendMessages,

	OwnerOnly: false,

	RunFunc: func(data *commands.HandlerData) {
		var dog DogAPI

		msg := data.Message
		err := queries.Get(api, &dog)
		if err != nil {
			msg.SendChannel("Error getting dog image")
			return
		}

		img := dog.Message
		msg.SendChannelEmbed(&discordgo.MessageEmbed{
			Title: "What the dog doin'",
			URL:   img,
			Image: &discordgo.MessageEmbedImage{
				URL: string(img),
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Powered by https://dog.ceo",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		})
	},
}
