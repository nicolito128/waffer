/*
	* Messages Package *

		The Messages package is responsible for providing
		common functionality for Message.Content manipulation.
*/
package messages

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var prefix = os.Getenv("BOT_PREFIX")

// Message
// provide a struct for common functionality
// for Message.Content manipulation.
type Message struct {
	Session *discordgo.Session
	MC      *discordgo.MessageCreate
	Content string
	Prefix  string
}

// New return a new Message structure.
func New(s *discordgo.Session, m *discordgo.MessageCreate) *Message {
	return &Message{Session: s, MC: m, Content: m.Content, Prefix: prefix}
}

// GetCommandWithPrefix return the command with the bot prefix: Ex: 'prefixCommandName'
func (msg *Message) GetCommandWithPrefix() string {
	args := strings.Split(msg.Content, " ")
	return args[0]
}

// GetCommand return the command without prefix: Ex: 'CommandName'
func (msg *Message) GetCommand() string {
	cmd := msg.GetCommandWithPrefix()
	cmdNoPrefix := strings.TrimPrefix(cmd, prefix)
	return cmdNoPrefix
}

// HasCommand check if the message content has command reference.
func (msg *Message) HasCommand() bool {
	command := msg.GetCommandWithPrefix()
	ok := strings.HasPrefix(msg.Content, command)
	if ok {
		return true
	}

	return false
}

// GetArguments return a []string of the command arguments. Command arguments are the
// inputs separated with an empty space in the message content.
func (msg *Message) GetArguments() []string {
	content := strings.TrimPrefix(msg.Content, msg.GetCommandWithPrefix())
	args := strings.Split(strings.TrimPrefix(content, " "), " ")
	return args
}

// GetPlainContent return the message content without prefix and command.
func (msg *Message) GetPlainContent() string {
	return strings.Join(msg.GetArguments(), " ")
}

// SendChannel sends a message to the message author channel.
func (msg *Message) SendChannel(str string, args ...any) {
	message := fmt.Sprintf(str, args...)
	msg.Session.ChannelMessageSend(msg.MC.ChannelID, message)
}

// SendChannelEmbed sends a message embed to the message author channel.
func (msg *Message) SendChannelEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return msg.Session.ChannelMessageSendEmbed(msg.MC.ChannelID, embed)
}

// HasHelPetition return a boolean if the message arguments ends with an "--help".
func (msg *Message) HasHelpPetition() bool {
	str := msg.GetPlainContent()
	ok := strings.HasSuffix(str, "--help")
	return ok
}
