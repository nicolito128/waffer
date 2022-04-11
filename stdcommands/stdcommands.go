package stdcommands

import (
	"github.com/nicolito128/waffer/plugins/commands"
	"github.com/nicolito128/waffer/stdcommands/dev"
	"github.com/nicolito128/waffer/stdcommands/ping"
)

type WafferCommand *commands.WafferCommand

// AddCommands load all the commands.
func AddCommands() {
	commands.AddRootCommands(
		ping.Command,
		dev.Command,
	)
}

func HasCommand(commandName string) bool {
	if commands.CommandList[commandName] != nil {
		return true
	}

	return false
}

func GetCommandList() commands.CList {
	return commands.CommandList
}
