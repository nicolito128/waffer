package github

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var link = "https://github.com/nicolito128/waffer"

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "github",
	Command: &plugins.CommandData{
		Description: "Github repository!",
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
	sm.ChannelSend("**Github repository**: %s", link)
}
