package github

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var repository = "https://github.com/nicolito128/waffer"

var Command = &commands.WafferCommand{
	Name:        "github",
	Aliases:     []string{"repository"},
	Description: "Code repository.",
	Category:    "development",

	Arguments:    []string{},
	RequiredArgs: 0,

	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		msg.SendChannel("**Github repository**: %s", repository)
	},
}
