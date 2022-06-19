package help

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "help",
	Aliases:     []string{"command"},
	Description: "Help for commands.",
	Category:    "info",

	Arguments:    []string{"<command>[optional]"},
	RequiredArgs: 0,

	OwnerOnly:          false,
	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		argument := strings.Join(msg.GetArguments(), " ")

		if argument == "" || argument == " " {
			msg.SendChannelSafe("You can enter the name of a command to get more information about it. Ex: `%shelp ping`.", msg.Prefix)
			msg.SendChannelSafe("You also have the option of entering **-h** or **--help** at the end of any command to get the information box. Ex: `%scalc -h` or `%sping --help`.", msg.Prefix, msg.Prefix)
			return
		}

		clist := commands.CommandList
		if clist[argument] == nil {
			msg.SendChannelSafe("Command **%s** not found.", argument)
			return
		}

		embed := commands.GetHelpEmbed(clist[argument])
		msg.SendChannelEmbed(embed)
	},
}
