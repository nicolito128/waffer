package commandlist

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "commandlist",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:        "commandlist",
		Description: "Shows a list of commands.",
		Category:    "info",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m.Message)
	clist := commands.CommandCollection

	var list []string
	for _, cmd := range clist {
		if strings.Contains(strings.Join(list, " "), cmd.Data.Name) {
			continue
		}

		list = append(list, cmd.Data.Name)
	}

	sm.ChannelSend("*List of commands* \n```\n"+strings.Join(list, ", ")+"```For more information about a command, type `%shelp <command>`.", sm.Prefix)
}
