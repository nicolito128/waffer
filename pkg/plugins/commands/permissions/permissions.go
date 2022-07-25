package permissions

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/config"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
)

// MemberHasPermission checks if a member has the given permission
func MemberHasPermission(s *discordgo.Session, m *discordgo.MessageCreate, cmd *commands.WafferCommand) bool {
	p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		return false
	}

	if p&int64(cmd.Data.Permissions.Require) == int64(cmd.Data.Permissions.Require) {
		return true
	}

	return false
}

// AllowDM returns true if the command comes from a DM and cmd.AllowDM is true.
func AllowDM(s *discordgo.Session, m *discordgo.MessageCreate, cmd *commands.WafferCommand) bool {
	ok, _ := ComesFromDM(s, m)
	if cmd.Data != nil {
		if ok && !cmd.Data.Permissions.AllowDM {
			return false
		}
	}

	return true
}

// ComesFromDM returns true if a message comes from a DM channel
func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}

// ValidPrefix returns true if the message starts with the prefix
func ValidPrefix(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	prefix := config.Config.Prefix

	if strings.HasPrefix(strings.Trim(m.Content, " "), prefix) {
		return true
	}

	return false
}

// Author verified if the message author is a bot or not
func ValidAuthor(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return false
	}

	// Ignore other bots
	if m.Author.Bot {
		return false
	}

	return true
}

// MessageHasArguments returns true if the message has the correct arguments length.
func MessageHasArguments(s *discordgo.Session, m *discordgo.MessageCreate, cmd *commands.WafferCommand) bool {
	if cmd.Data != nil {
		if cmd.Data.RequiredArgs > 0 {
			args := strings.Split(strings.Trim(m.Content, " "), " ")

			if len(args) < int(cmd.Data.RequiredArgs) {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You must provide more arguments. Use `%shelp %s` for more information.", config.Config.Prefix, cmd.Data.Name))
				return false
			}
		}
	}

	return true
}

func OwnerOnly(s *discordgo.Session, m *discordgo.MessageCreate, cmd *commands.WafferCommand) bool {
	if cmd.Data != nil {
		if cmd.Data.Permissions.OwnerOnly {
			if m.Author.ID == config.Config.OwnerID {
				return true
			} else {
				return false
			}
		}
	}

	return true
}

func HasHelpPetition(s *discordgo.Session, m *discordgo.MessageCreate, cmd *commands.WafferCommand) bool {
	if cmd.Data != nil {
		embed, err := cmd.HelpEmbed()
		if err != nil {
			return true
		}

		if strings.HasSuffix(m.Content, "--help") || strings.HasSuffix(m.Content, "-h") {
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
			return false
		}
	}

	return true
}
