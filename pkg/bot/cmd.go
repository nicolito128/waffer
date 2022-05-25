package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/utils/messages"
	"github.com/nicolito128/waffer/stdcommands"
)

func canMessageCommand(cmd stdcommands.WafferCommand, b *Bot, m *discordgo.MessageCreate, msg *messages.Message) bool {
	// DM check
	if m.GuildID == "" {
		// Command not allowed for dm channels
		if !cmd.RunInDM {
			msg.SendChannel("I can't use this command in direct messages.")
			return false
		}
	}

	// Help pettion in the command
	if msg.HasHelpPetition() {
		_, err := msg.SendChannelEmbed(getHelpEmbed(cmd))
		if err != nil {
			b.logs.Fatal(err.Error())
		}
		return false
	}

	// Arguments check
	if cmd.RequiredArgs > 0 {
		args := msg.GetArguments()
		if len(args) < cmd.RequiredArgs && args[0] == "" {
			msg.SendChannel("I need %d arguments for this command. Ask for help on this command using **%s --help**", cmd.RequiredArgs, msg.GetCommandWithPrefix())
			return false
		}
	}

	// Permissions check
	if cmd.DiscordPermissions > 0 {
		botPerms, err := b.session.State.UserChannelPermissions(b.session.State.User.ID, m.ChannelID)
		if err != nil {
			b.logs.Fatal(err.Error())
			return false
		}

		if (botPerms & cmd.DiscordPermissions) < cmd.DiscordPermissions {
			msg.SendChannel("Bot permissions are too low.")
			return false
		}

		perms, err := b.session.State.MessagePermissions(m.Message)
		if err != nil {
			b.logs.Fatal(err.Error())
			return false
		}

		if (perms & cmd.DiscordPermissions) < cmd.DiscordPermissions {
			msg.SendChannel("User permissions are too low.")
			return false
		}
	}

	return true
}

func getHelpEmbed(cmd stdcommands.WafferCommand) *discordgo.MessageEmbed {
	var args string
	if len(cmd.Arguments) == 0 {
		args = ""
	} else {
		args = "`" + strings.Join(cmd.Arguments, " ") + "`"
	}

	desc := fmt.Sprintf(`
		**Description**: %s
		**Aliases**: %s
		**Arguments**: %s
		**Category**: %s
	`,
		cmd.Description,
		strings.Join(cmd.Aliases, " "),
		args,
		""+cmd.Category+"",
	)

	return &discordgo.MessageEmbed{
		Title:       cmd.Name,
		Description: desc,
	}
}
