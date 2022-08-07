package invite

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
)

var link = "https://discord.com/oauth2/authorize?client_id=961017912355864617&scope=bot+applications.commands&permissions=1644401392759%20applications.commands"

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:        "invite",
		Type:        plugins.MessageCreateType,
		Handler:     Handler,
		Interaction: Interaction,
	},

	Data: &commands.CommandData{
		Name:        "invite",
		Description: "Shows the link to invite the bot to your server.",
		Category:    "info",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
		Slash: &discordgo.ApplicationCommand{
			Name:        "invite",
			Description: "Shows the link to invite the bot to your server.",
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Here is the invite link for your server: %s", link))
}

func Interaction(s *discordgo.Session, m *discordgo.InteractionCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Here is the invite link for your server: %s", link))
}
