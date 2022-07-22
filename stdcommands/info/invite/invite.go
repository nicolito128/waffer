package invite

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var link = "https://discord.com/oauth2/authorize?client_id=961017912355864617&scope=bot+applications.commands&permissions=8"

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "invite",
	Command: &plugins.CommandData{
		Description: "Invite me to your server!",
		Category:    "info",
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	sm.ChannelSend("Here is the invite link for your server: %s", link)
}
