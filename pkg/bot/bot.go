package bot

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/config"
)

/* New creates a new Gumdrop struct. */
func New(con *config.ConnectionConfig) (*discordgo.Session, error) {
	flag.StringVar(&con.Token, "token", con.Token, "Discord bot token")
	flag.StringVar(&con.OwnerID, "owner", con.OwnerID, "Discord bot owner")
	flag.StringVar(&con.Prefix, "prefix", con.Prefix, "Discord bot prefix")
	flag.Parse()

	config.Config = *con

	if con.Token == "" || len(con.Token) < 1 {
		return nil, config.ConfigTokenError
	}

	s, err := discordgo.New("Bot " + con.Token)
	if err != nil {
		return nil, err
	}

	s.Identify.Intents = con.Intents

	return s, nil
}

/* Init initializes the bot. */
func Init(g *discordgo.Session) error {
	err := g.Open()
	if err != nil {
		return err
	}

	fmt.Printf("Bot now running as %s. Press CTRL+C to exit.\n", g.State.User.Username)

	// Wait here until CTRL+C or other term signal is received.
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	// Close the connection
	return g.Close()
}
