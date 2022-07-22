package dog

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
	"github.com/nicolito128/waffer/pkg/queries"
)

type DogAPI struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

var api = "https://dog.ceo/api/breeds/image/random"

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "dog",
	Command: &plugins.CommandData{
		Description: "Sends a random dog image.",
		Category:    "fun",
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages & discordgo.PermissionAttachFiles,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var dog DogAPI
	sm := supermessage.New(s, m)

	err := queries.Get(api, &dog)
	if err != nil {
		sm.ChannelSend("Error getting dog image.")
		return
	}

	img := dog.Message
	sm.ChannelSendEmbed(&discordgo.MessageEmbed{
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
}
