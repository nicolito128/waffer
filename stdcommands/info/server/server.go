package server

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "server",
	Command: &plugins.CommandData{
		Description: "Server stats and info.",
		Permissions: plugins.CommandPermissions{
			AllowDM: false,
			Require: discordgo.PermissionSendMessages,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
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
		return
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

	sm.ChannelSendEmbed(&discordgo.MessageEmbed{
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
	})
}
