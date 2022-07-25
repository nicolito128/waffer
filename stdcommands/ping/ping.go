package ping

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "ping",
	Command: &plugins.CommandData{
		Description: "Ping!",
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! `%dms`", s.HeartbeatLatency().Milliseconds()))
}
