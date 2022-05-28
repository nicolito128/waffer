package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
	"github.com/nicolito128/waffer/plugins/utils/messages"
	"github.com/nicolito128/waffer/plugins/utils/permissions"
	"github.com/nicolito128/waffer/stdcommands"
)

func hasHelpPetition(cmd stdcommands.WafferCommand, b *Bot, msg *messages.Message) bool {
	if msg.HasHelpPetition() {
		_, err := msg.SendChannelEmbed(commands.GetHelpEmbed(cmd))
		if err != nil {
			b.logs.Fatal(err.Error())
		}
		return false
	}

	return true
}

func hasArgumentsForCommand(cmd stdcommands.WafferCommand, msg *messages.Message) bool {
	if cmd.RequiredArgs > 0 {
		args := msg.GetArguments()
		if len(args) < cmd.RequiredArgs || (len(args) == 1 && args[0] == "") {
			msg.SendChannel("I need %d arguments for this command. Ask for help at this command using **%s --help**", cmd.RequiredArgs, msg.GetCommandWithPrefix())
			return false
		}
	}

	return true
}

func canMessageCommand(cmd stdcommands.WafferCommand, b *Bot, m *discordgo.MessageCreate, msg *messages.Message) bool {
	// DM check
	dm, err := permissions.ComesFromDM(b.session, m)
	if err != nil {
		b.logs.Fatal(err.Error())
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
	helpPetition := hasHelpPetition(cmd, b, msg)
	if !helpPetition {
		return false
	}

	// Arguments check
	argCheck := hasArgumentsForCommand(cmd, msg)
	if !argCheck {
		return false
	}

	// Permissions check
	perms, err := permissions.MemberHasPermission(b.session, m.GuildID, m.Author.ID, cmd.DiscordPermissions)
	if err != nil {
		b.logs.Println(err.Error())
		return false
	}

	if !perms {
		msg.SendChannel("You don't have the permissions to use this command.")
		return false
	}

	return true
}
