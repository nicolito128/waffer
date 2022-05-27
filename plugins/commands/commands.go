package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/utils/messages"
)

// HandlerData is a struct for the RunFunc in commands.
// Provide common arguments for it.
type HandlerData struct {
	S       *discordgo.Session
	MC      *discordgo.MessageCreate
	Message *messages.Message
}

type RunHandler func(data *HandlerData)

type WafferCommand struct {
	Name        string   // Command name
	Aliases     []string // Command aliases. It's another way to invoke the same command
	Description string   // Command description for help.
	Category    string   // Group to which the command belongs.

	Arguments    []string // Arguments for the command.
	RequiredArgs int      // Number of arguments required for the command

	RunInDM bool // If the command can run in direct messages or not.

	DiscordPermissions int64 // Permissions required for the user and bot.

	RunFunc RunHandler // Run function
}

func (wc *WafferCommand) GetCategory() string {
	return wc.Category
}

func (wc *WafferCommand) GetDescription() string {
	return wc.Description
}

func (wc *WafferCommand) GetAliases() []string {
	return wc.Aliases
}

func (wc *WafferCommand) Args() ([]string, int) {
	return wc.Arguments, wc.RequiredArgs
}

// GetTrigger execute the current WafferCommand.RunFunc
func (wc *WafferCommand) GetTrigger(s *discordgo.Session, m *discordgo.MessageCreate, msg *messages.Message) {
	wc.RunFunc(&HandlerData{S: s, MC: m, Message: msg})
}

// GetHelpEmbed returns a discordgo.MessageEmbed with the help information of the command.
func GetHelpEmbed(cmd *WafferCommand) *discordgo.MessageEmbed {
	var args string
	if len(cmd.Arguments) == 0 {
		args = ""
	} else {
		args = "`" + strings.Join(cmd.Arguments, " ") + "`"
	}

	desc := fmt.Sprintf(`
		**Description**: %s
		**Aliases**: %s
		**Arguments**: %s
		**Category**: %s
	`,
		cmd.Description,
		strings.Join(cmd.Aliases, " "),
		args,
		""+cmd.Category+"",
	)

	return &discordgo.MessageEmbed{
		Title:       cmd.Name,
		Description: desc,
	}
}
