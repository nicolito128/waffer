package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
)

var CommandCollection = make(plugins.Collection[*WafferCommand])

type WafferCommand struct {
	Plugin *plugins.Plugin[*discordgo.MessageCreate]
	// If the plugin is a command sets some usefull information.
	Data *CommandData
}

// CommandData is a struct that contains some usefull information for a command.
type CommandData struct {
	// Name of the command.
	Name string
	// Command description
	Description string
	// Command arguments
	Arguments []string
	// Required arguments
	RequiredArgs uint
	// Command category
	Category string
	// Command permissions
	Permissions *CommandPermissions
}

type CommandPermissions struct {
	AllowDM   bool
	OwnerOnly bool
	Require   discordgo.PermissionOverwriteType
}

func (cmd *WafferCommand) HelpEmbed() (*discordgo.MessageEmbed, error) {
	if cmd.Data == nil {
		return nil, errors.New("Plugin is not a command.")
	}

	var desc string
	desc += fmt.Sprintf("**Description**: %s\n", cmd.Data.Description)

	if cmd.Data.Category != "" {
		desc += fmt.Sprintf("**Category**: `%s`\n", cmd.Data.Category)
	}

	if len(cmd.Data.Arguments) > 0 {
		desc += fmt.Sprintf("**Arguments**: \n`%s`", strings.Join(cmd.Data.Arguments, "\n"))
	}

	embed := &discordgo.MessageEmbed{
		Title:       cmd.Data.Name,
		Description: desc,
		Color:       0xe69b10,
	}
	return embed, nil
}

// Get returns a command from the CommandCollection if it exists.
func Get(name string) (*WafferCommand, error) {
	for _, cmd := range CommandCollection {
		if cmd.Data.Name == name {
			return cmd, nil
		}
	}

	return nil, fmt.Errorf("Command '%s' not found", name)
}

func AddList(commands ...*WafferCommand) {
	for _, cmd := range commands {
		Add(cmd)
	}
}

func Add(cmd *WafferCommand) error {
	if cmd.Data == nil {
		return errors.New("Command must have a CommandData struct.")
	}

	if cmd.Plugin.Name == "" && cmd.Data.Name == "" {
		return errors.New("Command must have a name.")
	}

	if len(cmd.Data.Name) == 0 || len(cmd.Plugin.Name) > 0 {
		fmt.Printf("Command '%s' added.\n", cmd.Plugin.Name)
		CommandCollection[cmd.Plugin.Name] = cmd
	} else {
		fmt.Printf("Command '%s' added.\n", cmd.Data.Name)
		CommandCollection[cmd.Data.Name] = cmd
	}

	return nil
}
