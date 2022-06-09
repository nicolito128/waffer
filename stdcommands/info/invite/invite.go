package invite

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var link = "https://discord.com/oauth2/authorize?client_id=961017912355864617&scope=bot+applications.commands&permissions=8"

var Command = &commands.WafferCommand{
	Name:        "invite",
	Aliases:     []string{},
	Description: "Get the invite link for your server.",
	Category:    "info",

	Arguments:    []string{},
	RequiredArgs: 0,

	OwnerOnly:          false,
	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		msg.SendChannel("Here is the invite link for your server: %s", link)
	},
}
