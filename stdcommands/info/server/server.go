package server

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:        "server",
		Type:        plugins.MessageCreateType,
		Handler:     Handler,
		Interaction: Interaction,
	},

	Data: &commands.CommandData{
		Name:        "server",
		Description: "Shows information about the server.",
		Category:    "info",
		Permissions: &commands.CommandPermissions{
			AllowDM: false,
			Require: discordgo.PermissionSendMessages,
		},
		Slash: &discordgo.ApplicationCommand{
			Name:        "server",
			Description: "Shows information about the server.",
		},
	},
}

func Server(s *discordgo.Session, m *discordgo.Message) *discordgo.MessageEmbed {
	sm := supermessage.New(s, m)
	var guild *discordgo.Guild

	for _, g := range s.State.Guilds {
		if g.ID == m.GuildID {
			guild = g
			break
		}
	}

	if guild == nil {
		sm.ChannelSendEmbed(&discordgo.MessageEmbed{
			Title:       "Error",
			Description: "Could not get guild channels.",
			Color:       0xFF0000,
		})
		return nil
	}

	var online, idle, dnd int
	for _, pres := range guild.Presences {
		if !pres.User.Bot {
			switch pres.Status {
			case discordgo.StatusOnline:
				online++
			case discordgo.StatusIdle:
				idle++
			case discordgo.StatusDoNotDisturb:
				dnd++
			}
		}
	}

	channels := len(guild.Channels)
	members := guild.MemberCount
	for _, member := range guild.Members {
		if member.User.Bot {
			members--
		}
	}

	return &discordgo.MessageEmbed{
		Title: guild.Name,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Channels", Value: fmt.Sprintf("%d", channels), Inline: true},
			{Name: "Users", Value: fmt.Sprintf("%d", members), Inline: true},
			{Name: "Presences", Value: fmt.Sprintf(":green_heart: %d\n:yellow_heart: %d\n:heart: %d", online, idle, dnd), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    fmt.Sprintf("Guild ID: %s", guild.ID),
			IconURL: guild.IconURL(),
		},
	}
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := Server(s, m.Message)
	if embed == nil {
		return
	}

	sm := supermessage.New(s, m.Message)
	sm.ChannelSendEmbed(embed)
}

func Interaction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := Server(s, i.Message)
	if embed == nil {
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Embeds:  []*discordgo.MessageEmbed{embed},
		},
	})
}
