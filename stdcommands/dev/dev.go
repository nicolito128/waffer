package dev

import (
	"github.com/nicolito128/waffer/plugins/commands"
)

var inviteURL = "https://discord.gg/yWqmnE4UmG"

var Command = &commands.WafferCommand{
	Name:        "dev",
	Aliases:     []string{"devserver"},
	Description: "Dev return the development bot server.",
	Category:    "development",

	Arguments:    []string{},
	RequiredArgs: 0,

	DiscordPermissions: 0,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message

		msg.SendChannel("**Development server**: %s", inviteURL)
	},
}
