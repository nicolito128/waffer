package commandlist

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "commandlist",
	Command: &plugins.CommandData{
		Description: "Get a list of all commands.",
		Category:    "info",
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	clist := plugins.CommandCollection

	var list []string
	for _, cmd := range clist {
		if strings.Contains(strings.Join(list, " "), cmd.Name) {
			continue
		}

		list = append(list, cmd.Name)
	}

	sm.ChannelSend("*List of commands* \n```\n"+strings.Join(list, ", ")+"```For more information about a command, type `%shelp <command>`.", sm.Prefix)
}
