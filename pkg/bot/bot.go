package bot

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/logger"
	"github.com/nicolito128/waffer/plugins/utils/checks"
	"github.com/nicolito128/waffer/plugins/utils/messages"
	"github.com/nicolito128/waffer/stdcommands"
)

var token = os.Getenv("BOT_TOKEN")
var prefix = os.Getenv("BOT_PREFIX")
var mode = os.Getenv("BOT_MODE")

// Bot
// provide a basic application struct.
type Bot struct {
	session *discordgo.Session   // Bot session
	logger  *logger.SystemLogger // Output information and errors
}

// New returns a new bot session and an error.
func New() (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Presence.Game.Name = prefix
	logger := logger.New(dg)

	bot := &Bot{dg, logger}
	return bot, nil
}

// Run open and starts the bot connection.
func (b *Bot) Run() {
	err := b.session.Open()
	if err != nil {
		b.logger.Fatalf("error opening connection, %s", err)
	}

	if mode == "debug" || mode == "" {
		b.logger.Println("Running debug mode.")
		go b.setStatusLog()
	}

	stdcommands.AddCommands()
	b.AddCommandsHandler()

	b.logger.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// AddCommandsHandler is in charge of setting the primary function
// that will check when users on any server try to use a command.
func (b *Bot) AddCommandsHandler() {
	b.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		clist := stdcommands.GetCommandList()
		for cmd, trigger := range clist {
			// Basic verification for no-bot attention.
			err := checks.Generals(s, m)
			if err != nil {
				break
			}

			msg := messages.New(s, m)

			if msg.GetCommand() == cmd.Name {
				ok := canMessageCommand(cmd, b, m, msg)
				if !ok {
					break
				}

				go trigger(s, m, msg)
				break
			}

			// If aliases exists
			if len(cmd.Aliases) >= 1 {
				for _, alias := range cmd.Aliases {
					if msg.GetCommand() == alias {
						ok := canMessageCommand(cmd, b, m, msg)
						if !ok {
							break
						}

						go trigger(s, m, msg)
						break
					}
				}
			}

		}
	})
}

// AddHandler set a new handler for the current session.
func (b *Bot) AddHandler(handler any) {
	b.session.AddHandler(handler)
}

// setStatusLog set console messages for debug mode, every 10 seconds.
func (b *Bot) setStatusLog() {
	tdr := time.Tick(10 * time.Second)
	for range tdr {
		b.logger.Printf(`Presence: %s | Guilds: %d | Message count: %d | Private channels: %d `,
			b.session.Identify.Presence.Game.Name,
			len(b.session.State.Guilds),
			len(b.session.State.PrivateChannels),
			b.session.State.MaxMessageCount)
	}
}
