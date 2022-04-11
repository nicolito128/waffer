package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/utils/messages"
)

type CList map[*WafferCommand]Trigger
type Trigger func(s *discordgo.Session, m *discordgo.MessageCreate, msg *messages.Message)

var CommandList = make(CList)

// AddRootCommands add in CommandList all the commands passed.
func AddRootCommands(cmds ...*WafferCommand) {
	for _, c := range cmds {
		CommandList[c] = c.GetTrigger
		fmt.Printf("Command '%s' added successfully!", c.Name)
	}
}
