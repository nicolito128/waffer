# Standard Commands
Standard Commands is a space to store and develop the different commands of the bot, which will be used by the users with the prefix provided to the application.

## Make a Command
The commands are stored in a folder that represents a single package, then these packages must be imported into the `stdcommands.go` file. This last file is in charge of loading the different commands when the bot is started.

Each command that is imported is a variable containing a WafferCommand struct, provided by the **commands** package, located in the `./plugins/commands` folder.

### Example
```go
    package excmd

    import (
        "github.com/nicolito128/waffer/plugins/commands"
    )

    var Command = &commands.WafferCommand{
	    Name:        "excmd",
	    Aliases:     []string{},
	    Description: "excmd makes stuff.",
	    Category:    "example",

	    Arguments:    []string{},
	    RequiredArgs: 0,

	    DiscordPermissions: 0,

	    RunInDM: true,

	    RunFunc: func(data *commands.HandlerData) {
            data.msg.SendChannel("Example command.")
        }
```
