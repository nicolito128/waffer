package stdcommands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/pkg/plugins"
)

type Middleware[T any] func(s *discordgo.Session, t T) bool
type CommandMiddleware func(s *discordgo.Session, t *discordgo.MessageCreate, c plugins.Plugin[*discordgo.MessageCreate]) bool

func Checker[T any](s *discordgo.Session, t T, mid ...Middleware[T]) error {
	for _, m := range mid {
		if !m(s, t) {
			return errors.New("Middleware failed.")
		}
	}

	return nil
}

func CommandChecker(s *discordgo.Session, t *discordgo.MessageCreate, cmd plugins.Plugin[*discordgo.MessageCreate], mid ...CommandMiddleware) error {
	for _, m := range mid {
		if !m(s, t, cmd) {
			return errors.New("Middleware failed.")
		}
	}

	return nil
}
