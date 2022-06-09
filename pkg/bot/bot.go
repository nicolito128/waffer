package bot

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/utils/checks"
	"github.com/nicolito128/waffer/plugins/utils/message_creation"
	"github.com/nicolito128/waffer/plugins/utils/messages"
	"github.com/nicolito128/waffer/stdcommands"
)

var token = os.Getenv("BOT_TOKEN")
var prefix = os.Getenv("BOT_PREFIX")
var mode = os.Getenv("BOT_MODE")

// Bot
// provide a basic application struct.
type Bot struct {
	session *discordgo.Session // Bot session
	logs    *log.Logger        // Output information and errors
}

// New returns a new bot session and an error.
func New() (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	dg.Identify.Presence.Game.Name = prefix
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentDirectMessageTyping | discordgo.IntentDirectMessages | discordgo.IntentGuildMessages | discordgo.IntentsMessageContent | discordgo.IntentGuildMembers | discordgo.IntentsAllWithoutPrivileged | discordgo.IntentGuildPresences | discordgo.Intent
	logs := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	bot := &Bot{dg, logs}
	return bot, nil
}

// Run open and starts the bot connection.
func (b *Bot) Run() {
	err := b.session.Open()
	if err != nil {
		b.logs.Fatalf("error opening connection, %s", err)
	}

	if mode == "debug" || mode == "" {
		b.logs.Println("Running debug mode.")
		go b.setStatusLog()
	}

	b.AddHandler(messageMentionBot)

	stdcommands.AddCommands()
	b.AddCommandsHandler()

	b.logs.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	b.session.Close()
}

// AddCommandsHandler is in charge of setting the primary function
// that will check when users on any server try to use a command.
func (b *Bot) AddCommandsHandler() {
	b.session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		clist := stdcommands.GetCommandList()
		msg := messages.New(s, m)
		commandName := msg.GetCommand()
		existsCommand := stdcommands.HasCommand(commandName)

		if existsCommand && strings.HasPrefix(msg.Content, msg.GetCommandWithPrefix()) {
			cmd := clist[commandName]

			// Basic verification for no-bot attention.
			err := checks.Generals(s, m)
			if err != nil {
				return
			}

			// If the command can be executed.
			ok := message_creation.UserCanRunCommand(cmd, s, m, msg)
			if !ok {
				return
			}

			// Execute the command in a new goroutine.
			go cmd.GetTrigger(s, m, msg)
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
		b.logs.Printf(`Presence: %s | Guilds: %d | Message count: %d | Private channels: %d `,
			b.session.Identify.Presence.Game.Name,
			len(b.session.State.Guilds),
			len(b.session.State.PrivateChannels),
			b.session.State.MaxMessageCount)
	}
}

// Every time someone mention the bot, this function will be called.
func messageMentionBot(s *discordgo.Session, m *discordgo.MessageCreate) {
	mentions := m.Mentions
	if s.State.User.Bot {
		return
	}

	if len(mentions) == 1 && m.Author.ID != s.State.User.ID && m.MessageReference == nil {
		for _, mention := range mentions {
			if mention.ID == s.State.User.ID {
				s.ChannelMessageSend(m.ChannelID, "Hello, I'm Waffer, a Discord bot made by @n128#5523. You can get more information about me with the commands `"+prefix+"help` and `"+prefix+"dev`.")
				return
			}
		}
	}
}
