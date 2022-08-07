package github

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var link = "https://github.com/nicolito128/waffer"

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:        "github",
		Type:        plugins.MessageCreateType,
		Handler:     Handler,
		Interaction: Interaction,
	},

	Data: &commands.CommandData{
		Name:        "github",
		Description: "Shows the link to the GitHub repository.",
		Category:    "info",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
		Slash: &discordgo.ApplicationCommand{
			Name:        "github",
			Description: "Shows the link to the GitHub repository.",
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m.Message)
	sm.ChannelSend("**Github repository**: %s", link)
}

func Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("**Github repository**: %s", link),
		},
	})
}
