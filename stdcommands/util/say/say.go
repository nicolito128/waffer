package say

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "say",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:         "say",
		Description:  "Says something with the bot.",
		Arguments:    []string{"<message>"},
		RequiredArgs: 1,
		Category:     "util",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
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
