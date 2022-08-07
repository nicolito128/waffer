package stdcommands

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/config"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/commands/permissions"
	"github.com/nicolito128/waffer/stdcommands/fun/animegirl"
	"github.com/nicolito128/waffer/stdcommands/fun/dog"
	"github.com/nicolito128/waffer/stdcommands/images/blur"
	"github.com/nicolito128/waffer/stdcommands/images/flip"
	"github.com/nicolito128/waffer/stdcommands/images/negative"
	"github.com/nicolito128/waffer/stdcommands/images/reflect"
	"github.com/nicolito128/waffer/stdcommands/info/avatar"
	"github.com/nicolito128/waffer/stdcommands/info/commandlist"
	"github.com/nicolito128/waffer/stdcommands/info/devserver"
	"github.com/nicolito128/waffer/stdcommands/info/github"
	"github.com/nicolito128/waffer/stdcommands/info/help"
	"github.com/nicolito128/waffer/stdcommands/info/invite"
	"github.com/nicolito128/waffer/stdcommands/info/server"
	"github.com/nicolito128/waffer/stdcommands/info/waffer"
	"github.com/nicolito128/waffer/stdcommands/ping"
	"github.com/nicolito128/waffer/stdcommands/util/calc"
	"github.com/nicolito128/waffer/stdcommands/util/say"
)

func LoadCommands(s *discordgo.Session) {
	commands.AddList(
		ping.Command,

		// Utils
		calc.Command,
		say.Command,

		// Information
		help.Command,
		commandlist.Command,
		avatar.Command,
		invite.Command,
		devserver.Command,
		github.Command,
		waffer.Command,
		server.Command,

		// APIs
		dog.Command,

		// Anime
		animegirl.Command,

		// Images
		flip.Command,
		negative.Command,
		reflect.Command,
		blur.Command,
	)

	s.AddHandler(Command)
	s.AddHandler(LoadInteraction)
	s.AddHandler(ExecuteInteraction)
}

func Command(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := Checker(s, m, permissions.ValidPrefix, permissions.ValidAuthor)
	if err != nil {
		return
	}

	cmd := strings.Split(strings.Replace(m.Content, config.Config.Prefix, "", 1), " ")[0]
	command, err := commands.Get(cmd)
	if err != nil {
		return
	}

	if err = CommandChecker(s, m, command,
		permissions.AllowDM,
		permissions.MessageHasArguments,
		permissions.OwnerOnly,
		permissions.HasHelpPetition,
		permissions.MemberHasPermission,
	); err != nil {
		return
	}
	go command.Plugin.Handler(s, m)
}

func LoadInteraction(s *discordgo.Session, g *discordgo.GuildCreate) {
	for _, c := range commands.CommandCollection {
		if c.Data != nil && c.Data.Slash != nil {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, g.ID, c.Data.Slash)
			if err != nil {
				log.Panicf("Error creating interaction: %s", err.Error())
			}

			fmt.Println("Loaded interaction:", c.Data.Slash.Name)
		}
	}

}

func ExecuteInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	cmd, err := commands.Get(i.ApplicationCommandData().Name)
	if err != nil {
		return
	}

	go cmd.Plugin.Interaction(s, i)
}
