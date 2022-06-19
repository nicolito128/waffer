package commandlist

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "commandlist",
	Aliases:     []string{"commands"},
	Description: "Get a list of commands.",
	Category:    "info",

	Arguments:    []string{},
	RequiredArgs: 0,

	OwnerOnly:          false,
	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		clist := commands.CommandList

		var list []string
		for _, cmd := range clist {
			if strings.Contains(strings.Join(list, " "), cmd.Name) {
				continue
			}
			list = append(list, cmd.Name)
		}

		msg.SendChannel("*List of commands*")
		msg.SendChannel("```\n" + strings.Join(list, ", ") + "```")
		msg.SendChannel("For more information about a command, type `%shelp <command>`.", msg.Prefix)
	},
}
