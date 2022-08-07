package avatar

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

// Create commands.WafferCommand for avatar command
var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "avatar",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:        "avatar",
		Description: "Returns an embed with the user avatar profile picture. Optional you can mention one member.",
		Category:    "info",
		Arguments:   []string{"[@user]"},
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var avatarURL string
	var user *discordgo.User
	sm := supermessage.New(s, m.Message)

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
