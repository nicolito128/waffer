package avatar

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "avatar",
	Command: &plugins.CommandData{
		Description: "Returns an embed with the user avatar profile picture. Optional you can mention one member.",
		Arguments:   []string{"<mention[optional]>"},
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var avatarURL string
	var user *discordgo.User
	sm := supermessage.New(s, m)

	if len(m.Mentions) >= 1 {
		avatarURL = m.Mentions[0].AvatarURL("500x500")
		user = m.Mentions[0]
	} else {
		avatarURL = m.Author.AvatarURL("500x500")
		user = m.Author
	}

	sm.ChannelSendEmbed(&discordgo.MessageEmbed{
		Title: user.Username + "'s avatar",
		URL:   avatarURL,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    user.Username,
			IconURL: avatarURL,
		},
		Image: &discordgo.MessageEmbedImage{
			URL:    avatarURL,
			Width:  500,
			Height: 500,
		},
	})
}
