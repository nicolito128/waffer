package devserver

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var link = "https://discord.gg/yWqmnE4UmG"

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "devserver",
	Command: &plugins.CommandData{
		Description: "Development discord server.",
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
	sm.ChannelSend("**Development server**: %s", link)
}
