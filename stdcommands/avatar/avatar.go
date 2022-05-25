package avatar

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "avatar",
	Aliases:     []string{"avy"},
	Description: "Avatar return an embed with the user avatar profile picture. Optional you can mention one member.",
	Category:    "info",

	Arguments:    []string{"mention[optional]"},
	RequiredArgs: 0,

	DiscordPermissions: 0,

	RunInDM: false,

	RunFunc: func(data *commands.HandlerData) {
		var avatarURL string
		var user *discordgo.User
		msg := data.Message
		mentions := data.MC.Mentions

		if len(mentions) >= 1 {
			avatarURL = mentions[0].AvatarURL("500x500")
			user = mentions[0]
		} else {
			avatarURL = data.MC.Author.AvatarURL("500x500")
			user = data.MC.Author
		}

		msg.SendChannelEmbed(&discordgo.MessageEmbed{
			Title: user.Username + "'s avatar",
			URL:   avatarURL,
			Color: user.AccentColor,
			Image: &discordgo.MessageEmbedImage{
				URL:    avatarURL,
				Width:  500,
				Height: 500,
			},
		})
	},
}
