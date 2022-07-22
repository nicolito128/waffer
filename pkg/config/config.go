package config

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Config ConnectionConfig

// ConnectionConfig to connect the bot and other features.
type ConnectionConfig struct {
	// Bot Token
	Token string `json:"token"`
	// Bot Intents
	Intents discordgo.Intent `json:"intents"`
	// Bot Onwer ID
	OwnerID string `json:"owner"`
	// Bot Prefix
	Prefix string `json:"prefix"`
}

type ConfigError struct {
	Flag    string
	Message string
}

func (ge *ConfigError) Error() string {
	return fmt.Sprintf("Config error by flag '%s': %s", ge.Flag, ge.Message)
}

var ConfigTokenError = &ConfigError{"token", "Token is required."}
