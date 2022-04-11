package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/utils/messages"
)

type CList map[string]*WafferCommand
type Trigger func(s *discordgo.Session, m *discordgo.MessageCreate, msg *messages.Message)

var CommandList = make(CList)

// AddRootCommands add in CommandList all the commands passed.
func AddRootCommands(cmds ...*WafferCommand) {
	for _, c := range cmds {
		CommandList[c.Name] = c
		fmt.Printf("Command '%s' added successfully!\n", c.Name)

		if len(c.Aliases) >= 1 {
			for _, alias := range c.Aliases {
				CommandList[alias] = c
				fmt.Printf("Command alias '%s' added successfully!\n", alias)
			}
		}
	}
}
