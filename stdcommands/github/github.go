package github

import "github.com/nicolito128/waffer/plugins/commands"

var repository = "https://github.com/nicolito128/waffer"

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
		msg.SendChannel("**Github repository**: %s", repository)
	},
}
