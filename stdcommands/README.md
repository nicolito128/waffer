# Standard Commands
Standard Commands is a space to store and develop the different commands of the bot, which will be used by the users with the prefix provided to the application.

## Make a Command
The commands are stored in a folder that represents a single package, then these packages must be imported into the stdcommands.go file. Stdcommands load the different commands when the bot is started.

Each command is imported as a variable containing a generic `Plugin[*discordgo.MessageCreate]` struct, provided by the plugins package located at `./pkg/plugins/plugins.go`.

## Example
```go
package ping

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
)

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "ping",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:        "ping",
		Description: "Ping!",
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

```