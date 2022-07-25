package invite

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var link = "https://discord.com/oauth2/authorize?client_id=961017912355864617&scope=bot+applications.commands&permissions=1644905757943"

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "invite",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:        "invite",
		Description: "Shows the link to invite the bot to your server.",
		Category:    "info",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	sm.ChannelSend("Here is the invite link for your server: %s", link)
}
