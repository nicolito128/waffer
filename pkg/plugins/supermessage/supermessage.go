package supermessage

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/config"
)

// SuperMessage
// provide a struct for common functionality
// for Message.Content manipulation.
type SuperMessage struct {
	Self    *discordgo.Message
	Content string
	Session *discordgo.Session
	Create  *discordgo.MessageCreate
	Prefix  string
}

var prefix = config.Config.Prefix

// New return a new Message structure.
func New(s *discordgo.Session, m *discordgo.MessageCreate) *SuperMessage {
	return &SuperMessage{
		Self:    m.Message,
		Content: m.Content,
		Session: s,
		Create:  m,
		Prefix:  prefix,
	}
}

// Command returns the first element who starts with the prefix.
func (sm *SuperMessage) Command() string {
	return strings.Split(strings.Trim(sm.Content, " "), " ")[0]
}

// StartsWithPrefix returns true if the message starts with the prefix.
func (sm *SuperMessage) StartsWithPrefix(str string) bool {
	return strings.HasPrefix(strings.Trim(sm.Content, " "), prefix)
}

// Arguments returns a []string without the prefix!command element.
func (sm *SuperMessage) Arguments() []string {
	content := strings.Trim(strings.TrimPrefix(strings.Trim(sm.Content, " "), sm.Command()), " ")
	args := strings.Split(content, " ")
	if len(args) >= 2 {
		args = args[1:]
		println(args)
	} else {
		args = []string{}
	}

	return args
}

// PlainContent returns the message content without the prefix!command.
func (sm *SuperMessage) PlainContent() string {
	return strings.Trim(strings.Join(sm.Arguments(), " "), " ")
}

// ChannelSend sends a message to the message author channel without mentions allowed.
func (sm *SuperMessage) ChannelSend(str string, args ...any) (*discordgo.Message, error) {
	message := fmt.Sprintf(str, args...)
	m, err := sm.Session.ChannelMessageSendComplex(sm.Create.ChannelID, &discordgo.MessageSend{
		Content: message,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{},
			Roles: []string{},
		},
	})
	return m, err
}

func (sm *SuperMessage) ChannelSendEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	if embed.Color == 0 {
		embed.Color = sm.Session.State.MessageColor(sm.Create.Message)
	}

	m, err := sm.Session.ChannelMessageSendComplex(sm.Create.ChannelID, &discordgo.MessageSend{
		Embed: embed,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{},
			Roles: []string{},
		},
	})

	return m, err
}

func (sm *SuperMessage) ChannelSendEmbedUnsafe(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	if embed.Color == 0 {
		embed.Color = sm.Session.State.MessageColor(sm.Create.Message)
	}

	return sm.Session.ChannelMessageSendEmbed(sm.Create.ChannelID, embed)
}

// ChannelSendUnsafe sends a message to the message author channel with mentions allowed.
func (sm *SuperMessage) ChannelSendUnsafe(str string, args ...any) (*discordgo.Message, error) {
	message := fmt.Sprintf(str, args...)
	m, err := sm.Session.ChannelMessageSend(sm.Create.ChannelID, message)
	return m, err
}

// ChannelSendComplex send a message with complex options.
func (sm *SuperMessage) ChannelSendComplex(data *discordgo.MessageSend) (*discordgo.Message, error) {
	m, err := sm.Session.ChannelMessageSendComplex(sm.Create.ChannelID, data)
	return m, err
}
