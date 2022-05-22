package stdcommands

import (
	"github.com/nicolito128/waffer/plugins/commands"
	"github.com/nicolito128/waffer/stdcommands/avatar"
	"github.com/nicolito128/waffer/stdcommands/calc"
	"github.com/nicolito128/waffer/stdcommands/dev"
	"github.com/nicolito128/waffer/stdcommands/ping"
)

type WafferCommand *commands.WafferCommand

// AddCommands load all the commands.
func AddCommands() {
	commands.AddRootCommands(
		ping.Command,
		dev.Command,
		avatar.Command,
		calc.Command,
	)
}

// HasCommand return true if command list has the command name passed.
func HasCommand(commandName string) bool {
	if commands.CommandList[commandName] != nil {
		return true
	}

	return false
}

// GetCommandList return a map[string]:*WafferCommand
func GetCommandList() commands.CList {
	return commands.CommandList
}
