/*
	* Checks Package *

		The Checks package is intended to provide functionality
		to verify core properties of users before they execute
		any command.
*/
package checks

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/env"
	"github.com/nicolito128/waffer/plugins/utils/messages"
)

var prefix = env.GetBotPrefix()

// Generals check if all the basic verification are successfully passed.
func Generals(s *discordgo.Session, m *discordgo.MessageCreate) (err error) {
	err = Prefix(m)
	if err != nil {
		return err
	}

	err = Author(s, m)
	if err != nil {
		return err
	}

	return nil
}

// Prefix verified if the message content starts with de bot prefix
func Prefix(m *discordgo.MessageCreate) error {
	msg := messages.New(nil, m)
	cmd := msg.GetCommandWithPrefix()
	content := m.Content

	if ok := strings.HasPrefix(content, cmd); ok {
		return nil
	}

	return errors.New("No prefix.")
}

// Author verified if the message author is a bot or not
func Author(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return errors.New("Auto invocation not allowed.")
	}

	// Ignore other bots
	if m.Author.Bot {
		return errors.New("No bots allowed.")
	}

	return nil
}
