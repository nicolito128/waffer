package help

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "help",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:         "help",
		Description:  "Gets help about a command.",
		Category:     "info",
		Arguments:    []string{"<command>"},
		RequiredArgs: 1,
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m.Message)
	args := strings.Trim(sm.PlainContent(), " ")

	cmd, err := commands.Get(args)
	if err != nil {
		sm.ChannelSend("You can enter the name of a command to get more information about it. Ex: `%shelp ping`.\nYou also have the option of entering **-h** or **--help** at the end of any command to get the information box. Ex: `%scalc -h` or `%sping --help", sm.Prefix, sm.Prefix, sm.Prefix)
		return
	}

	embed, err := cmd.HelpEmbed()
	if err != nil {
		return
	}

	sm.ChannelSendEmbed(embed)
}
