package rps

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/waffer/plugins/commands"
)

var moves = []string{"rock", "paper", "scissors"}

func applyPriority(firstMove, secondMove string) bool {
	first := parseShortMove(firstMove)
	second := parseShortMove(secondMove)

	if first == "rock" && second == "scissors" {
		return true
	}

	if first == "paper" && second == "rock" {
		return true
	}

	if first == "scissors" && second == "paper" {
		return true
	}

	return false
}

func parseShortMove(move string) string {
	switch move {
	case "r":
		return "rock"
	case "p":
		return "paper"
	case "s":
		return "scissors"
	default:
		return move
	}
}

func getEmoji(move string) string {
	switch move {
	case "rock", "r":
		return ":fist:"
	case "paper", "p":
		return ":raised_back_of_hand:"
	case "scissors", "s":
		return ":vulcan:"
	}

	return ""
}

var Command = &commands.WafferCommand{
	Name:        "rps",
	Aliases:     []string{"rockpaperscissors"},
	Description: "Play rock paper scissors with the bot.",
	Category:    "fun",

	Arguments:    []string{"rock | paper | scissors"},
	RequiredArgs: 1,

	OwnerOnly:          false,
	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		content := strings.ToLower(strings.Join(msg.GetArguments(), " "))

		move := parseShortMove(content)
		if !strings.Contains(strings.Join(moves, " "), move) {
			msg.SendChannel("You need to specify a move like rock (r), paper (p) or scissors (s).")
			return
		}

		rand.Seed(time.Now().UnixNano())
		rbInd := rand.Intn(len(moves[0:3]))
		rbMove := moves[rbInd]

		// Tie case
		if move == rbMove {
			msg.SendChannel("I use %s (%s) and you use %s (%s).", getEmoji(rbMove), rbMove, getEmoji(move), move)
			msg.SendChannel("We tied! :handshake:")
			return
		}

		// If result is true then the first move is the winner.
		// If result is false then the second move is the winner.
		result := applyPriority(content, rbMove)
		if result {
			msg.SendChannel("I use %s (%s) and you use %s (%s). You win!", getEmoji(rbMove), rbMove, getEmoji(move), move)
			msg.SendChannel(":confetti_ball: :tada: :tada: :tada: :confetti_ball:")
		} else {
			msg.SendChannel("I use %s (%s) and you use %s (%s). I win!", getEmoji(rbMove), rbMove, getEmoji(move), move)
			msg.SendChannel("You lost! :rofl:")
		}
	},
}
