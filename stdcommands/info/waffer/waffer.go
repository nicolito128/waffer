package waffer

import (
	"fmt"
	"runtime"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "waffer",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:        "waffer",
		Description: "Shows Waffer information.",
		Category:    "info",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
		Slash: &discordgo.ApplicationCommand{
			Name:        "waffer",
			Description: "Shows Waffer information.",
		},
	},
}

func Waffer(s *discordgo.Session, m *discordgo.Message) {
	sm := supermessage.New(s, m)
	var channels, guilds, users int
	guilds = len(s.State.Guilds)

	for _, guild := range s.State.Guilds {
		channels += len(guild.Channels)

		for _, member := range guild.Members {
			if !member.User.Bot {
				users++
			}
		}
	}

	sm.ChannelSendEmbed(&discordgo.MessageEmbed{
		Title:       "Waffer",
		Description: "Information about me and how many things I'm doing.",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Guilds", Value: fmt.Sprintf("%d", guilds), Inline: true},
			{Name: "Channels", Value: fmt.Sprintf("%d", channels), Inline: true},
			{Name: "Members", Value: fmt.Sprintf("%d", users), Inline: true},
			{Name: "Commands", Value: fmt.Sprintf("%d", len(commands.CommandCollection)), Inline: true},
			{Name: "Ping", Value: fmt.Sprintf("%dms", s.HeartbeatLatency().Milliseconds()), Inline: true},
			{Name: "OS", Value: runtime.GOOS, Inline: true},
			{Name: "Go Version", Value: runtime.Version(), Inline: true},
			{Name: "Goroutines", Value: fmt.Sprintf("%d", runtime.NumGoroutine()), Inline: true},
			{Name: "CPU Available", Value: fmt.Sprintf("%d", runtime.NumCPU()), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Repository: https://github.com/nicolito128/waffer",
			IconURL: "https://i.imgur.com/TXjXenF.png",
		},
	})
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	Waffer(s, m.Message)
}

func Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	Waffer(s, i.Message)
}
