package say

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "say",
	Aliases:     []string{"speak", "speakup"},
	Description: "Make the bot say something.",
	Category:    "util",

	Arguments:    []string{"<content>"},
	RequiredArgs: 1,

	OwnerOnly:          false,
	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: false,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		content := strings.Join(msg.GetArguments(), " ")

		if content == "" || content == " " {
			msg.SendChannel("You need to specify something to say.")
			return
		}

		mentions := len(data.MC.MentionRoles)
		if data.MC.MentionEveryone || (mentions > 0) {
			msg.SendChannel("You can't mention everyone or roles in a say command.")
			return
		}

		data.S.ChannelMessageDelete(data.MC.ChannelID, data.MC.ID)
		msg.SendChannel(content)
	},
}
