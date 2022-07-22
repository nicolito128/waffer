package say

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "say",
	Command: &plugins.CommandData{
		Description:  "Says something with the bot.",
		Arguments:    []string{"<message>"},
		RequiredArgs: 1,
		Category:     "util",
		Permissions: plugins.CommandPermissions{
			AllowDM: false,
			Require: discordgo.PermissionSendMessages,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	content := sm.PlainContent()

	if strings.Trim(content, " ") == "" || len(content) < 1 {
		sm.ChannelSend("You need to specify something to say.")
		return
	}

	if len(content) > 2000 {
		sm.ChannelSend("The message is too long.")
		return
	}

	sm.ChannelSend(content)
}
