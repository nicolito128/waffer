package ping

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "ping",
	Aliases:     []string{"pong"},
	Description: "Ping sends the current bot latency in milliseconds.",
	Category:    "info",

	Arguments:    []string{},
	RequiredArgs: 0,

	RunInDM: true,

	DiscordPermissions: discordgo.PermissionSendMessages,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		latency := data.S.HeartbeatLatency().Milliseconds()

		if msg.GetCommand() == "pong" {
			msg.SendChannel("Ping! Latency: %dms", latency)
			return
		}

		msg.SendChannel("Pong! Latency: %dms", latency)
	},
}
