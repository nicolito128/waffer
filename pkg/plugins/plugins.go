package plugins

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Collection[T any] map[string]T

type PluginType string

const (
	MessageCreateType PluginType = "message_create"
)

// PluginHandler represents the structs that will be used to handle the plugins.
// Like discordgo.Session.
type PluginHandler interface {
	AddHandler(interface{}) func()
	AddHandlerOnce(interface{}) func()
}

// Plugin represents a plugin for the bot.
type Plugin[T any] struct {
	// Plugin name.
	Name string
	// Plugin type.
	Type PluginType
	// Handler is the function that will be called for the differents plugins
	Handler func(*discordgo.Session, T)
}

// Adds a plugin to the collection.
func Add[T any](col Collection[*Plugin[T]], plugin ...*Plugin[T]) {
	for _, p := range plugin {
		fmt.Printf("Plugin '%s' added.\n", p.Name)
		col[p.Name] = p
	}
}

// Gets a plugin from the collection.
func Get[T any](col Collection[*Plugin[T]], key string) (*Plugin[T], error) {
	for _, plugin := range col {
		if plugin.Name == key {
			return plugin, nil
		}
	}

	return nil, fmt.Errorf("Plugin '%s' not found", key)
}
