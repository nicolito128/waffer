# Standard Commands
Standard Commands is a space to store and develop the different commands of the bot, which will be used by the users with the prefix provided to the application.

## Make a Command
The commands are stored in a folder that represents a single package, then these packages must be imported into the `stdcommands.go` file. Stdcommands load the different commands when the bot is started.

Each command is imported as a variable containing a `WafferCommand` struct, provided by the **commands** package located at `./plugins/commands`.

### Example
```go
    package excmd

    import (
		"github.com/bwmarrin/discordgo"
        "github.com/nicolito128/waffer/plugins/commands"
    )

    var Command = &commands.WafferCommand{
	    Name:        "excmd",
	    Aliases:     []string{},
	    Description: "excmd makes stuff.",
	    Category:    "example",

	    Arguments:    []string{},
	    RequiredArgs: 0,

	    DiscordPermissions: discordgo.PermissionSendMessages,

	    RunInDM: true,

	    RunFunc: func(data *commands.HandlerData) {
            data.msg.SendChannel("Example command.")
        }
```
