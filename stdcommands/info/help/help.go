package help

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var prefix = os.Getenv("BOT_PREFIX")

var Command = &commands.WafferCommand{
	Name:        "help",
	Aliases:     []string{"command"},
	Description: "Help for commands.",
	Category:    "info",

	Arguments:    []string{"<command>[optional]"},
	RequiredArgs: 0,

	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		argument := strings.Join(msg.GetArguments(), " ")

		if argument == "" || argument == " " {
			msg.SendChannel("You can enter the name of a command to get more information about it. Ex: `%shelp ping`.", prefix)
			msg.SendChannel("You also have the option of entering **-h** or **--help** at the end of any command to get the information box. Ex: `%scalc -h` or `%sping --help`.", prefix, prefix)
			return
		}

		clist := commands.CommandList
		if clist[argument] == nil {
			msg.SendChannel("Command **%s** not found.", argument)
			return
		}

		embed := commands.GetHelpEmbed(clist[argument])
		msg.SendChannelEmbed(embed)
	},
}
