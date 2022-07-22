package animegirl

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	animegirls "github.com/nicolito128/animegirls-holding-programming-books"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "animegirl",
	Command: &plugins.CommandData{
		Description: "Sends a random anime holding a programming book. Based repository: https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books",
		Category:    "fun",
		Arguments:   []string{"<programming language>[optional]"},
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages & discordgo.PermissionAttachFiles,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	content := sm.PlainContent()

	if content == "" || content == " " {
		rbIndex := rand.Intn(len(animegirls.Languages))
		rbLang := animegirls.Languages[rbIndex]

		im, err := animegirls.GetRandomImage(rbLang)
		if err != nil {
			sm.ChannelSend("Error: %s", err.Error())
		}

		sm.ChannelSendEmbed(&discordgo.MessageEmbed{
			URL: im,
			Image: &discordgo.MessageEmbedImage{
				URL: im,
			},
		})
		return
	}

	im, err := animegirls.GetRandomImage(content)
	if err != nil && im == "" {
		sm.ChannelSend("No images found for that language.")
		sm.ChannelSend(fmt.Sprintf("Try one of these languages: `%s`", strings.Join(animegirls.Languages, ", ")))
		return
	}

	sm.ChannelSendEmbed(&discordgo.MessageEmbed{
		URL: im,
		Image: &discordgo.MessageEmbedImage{
			URL: im,
		},
	})
}
