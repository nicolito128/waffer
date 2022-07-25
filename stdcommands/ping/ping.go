package ping

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "ping",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:        "ping",
		Description: "Ping!",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! `%dms`", s.HeartbeatLatency().Milliseconds()))
}
