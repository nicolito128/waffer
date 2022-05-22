package calc

import (
	"strings"

	"github.com/nicolito128/go-calculator"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "calculator",
	Aliases:     []string{"calc"},
	Description: "For mathematics.",
	Category:    "science",

	Arguments:    []string{},
	RequiredArgs: 0,

	DiscordPermissions: 0,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		args := msg.GetArguments()

		var decimals string
		if len(args) >= 2 {
			if strings.HasSuffix(args[0], "f") {
				decimals = args[0]
				args = args[1:]
			}
		}

		content := strings.Join(args, "")
		res, err := calculator.Resolve(content)
		if err != nil {
			msg.SendChannel("Error: %s", err)
			return
		}

		switch decimals {
		case "0f":
			msg.SendChannel("%s = %.0f", content, res)
		case "3f":
			msg.SendChannel("%s = %.3f", content, res)
		case "4f":
			msg.SendChannel("%s = %.4f", content, res)
		case "5f":
			msg.SendChannel("%s = %.5f", content, res)
		case "6f":
			msg.SendChannel("%s = %.6f", content, res)
		case "f":
			msg.SendChannel("%s = %f", content, res)
		default:
			msg.SendChannel("%s = %.2f", content, res)
		}
	},
}
