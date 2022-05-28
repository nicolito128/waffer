package message_creation

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
	"github.com/nicolito128/waffer/plugins/utils/messages"
	"github.com/nicolito128/waffer/plugins/utils/permissions"
	"github.com/nicolito128/waffer/stdcommands"
)

func MessageHasHelpPetition(cmd stdcommands.WafferCommand, msg *messages.Message) bool {
	if msg.HasHelpPetition() {
		msg.SendChannelEmbed(commands.GetHelpEmbed(cmd))
		return false
	}

	return true
}

func MessageHasRequiredArguments(cmd stdcommands.WafferCommand, msg *messages.Message) bool {
	if cmd.RequiredArgs > 0 {
		args := msg.GetArguments()
		if len(args) < cmd.RequiredArgs || (len(args) == 1 && args[0] == "") {
			return false
		}
	}

	return true
}

func UserCanRunCommand(cmd stdcommands.WafferCommand, s *discordgo.Session, m *discordgo.MessageCreate, msg *messages.Message) bool {
	// DM check
	dm, err := permissions.ComesFromDM(s, m)
	if err != nil {
		return false
	}

	if dm {
		if cmd.RunInDM {
			return true
		} else {
			msg.SendChannel("This command can't be used in DM.")
			return false
		}
	}

	// Help pettion in the command
	helpPetition := MessageHasHelpPetition(cmd, msg)
	if !helpPetition {
		return false
	}

	// Arguments check
	argCheck := MessageHasRequiredArguments(cmd, msg)
	if !argCheck {
		msg.SendChannel("I need %d arguments for this command. Ask for help at this command using `%s --help`", cmd.RequiredArgs, msg.GetCommandWithPrefix())
		return false
	}

	// Permissions check
	perms, err := permissions.MemberHasPermission(s, m.GuildID, m.Author.ID, cmd.DiscordPermissions)
	if err != nil {
		return false
	}

	if !perms {
		msg.SendChannel("You don't have the permissions to use this command.")
		return false
	}

	return true
}
