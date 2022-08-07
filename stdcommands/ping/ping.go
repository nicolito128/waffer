package ping

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:        "ping",
		Type:        plugins.MessageCreateType,
		Handler:     Handler,
		Interaction: Interaction,
	},

	Data: &commands.CommandData{
		Name:        "ping",
		Description: "Ping!",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
		Slash: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Waffer slash ping",
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Pong! `%dms`", s.HeartbeatLatency().Milliseconds()))
}

func Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Pong! `%dms`", s.HeartbeatLatency().Milliseconds()),
		},
	})
}
