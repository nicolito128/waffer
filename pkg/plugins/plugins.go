package plugins

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var CommandCollection = make(Collection[Plugin[*discordgo.MessageCreate]])

type Collection[T any] map[string]T

// PluginHandler represents the structs that will be used to handle the plugins.
// Like Gumdrop or discordgo.Session.
type PluginHandler interface {
	AddHandler(interface{}) func()
	AddHandlerOnce(interface{}) func()
}

// GumdropPlugin represents a plugin for the grumdrop bot.
type Plugin[T any] struct {
	// Plugin name.
	Name string
	// If the plugin is a command sets some usefull information.
	Command *CommandData
	// Handler is the function that will be called for the differents plugins
	Handler func(*discordgo.Session, T)
}

// CommandData is a struct that contains some usefull information for a command.
type CommandData struct {
	// Command description
	Description string
	// Command arguments
	Arguments []string
	// Required arguments
	RequiredArgs uint
	// Command category
	Category string
	// Command permissions
	Permissions CommandPermissions
}

type CommandPermissions struct {
	AllowDM   bool
	OwnerOnly bool
	Require   discordgo.PermissionOverwriteType
}

func (p Plugin[T]) HelpEmbed() (*discordgo.MessageEmbed, error) {
	if p.Command == nil {
		return nil, errors.New("Plugin is not a command.")
	}

	var desc string
	desc += fmt.Sprintf("**Description**: %s\n", p.Command.Description)

	if p.Command.Category != "" {
		desc += fmt.Sprintf("**Category**: `%s`\n", p.Command.Category)
	}

	if len(p.Command.Arguments) > 0 {
		desc += fmt.Sprintf("**Arguments**: `%s`\n", strings.Join(p.Command.Arguments, ", "))
	}

	embed := &discordgo.MessageEmbed{
		Title:       p.Name,
		Description: desc,
		Color:       0xe69b10,
	}
	return embed, nil
}

// AddPlugin adds a plugin to the grumdrop PluginCollection.
func AddPlugin[T any](col Collection[Plugin[T]], plugin ...Plugin[T]) {
	for _, p := range plugin {
		fmt.Printf("Plugin '%s' added.\n", p.Name)
		col[p.Name] = p
	}
}

func GetPlugin[T any](col Collection[Plugin[T]], key string) (Plugin[T], error) {
	for _, plugin := range col {
		if plugin.Name == key {
			return plugin, nil
		}
	}

	return col[""], fmt.Errorf("Plugin '%s' not found", key)
}
