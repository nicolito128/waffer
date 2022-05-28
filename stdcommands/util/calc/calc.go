package calc

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/go-calculator"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "calculator",
	Aliases:     []string{"calc"},
	Description: "A mathematical calculation command. Use the example: `waf!calculator (10 / 2) + 1`, `waf!calculator 4f (1 / 9)`. The argument `<number>f` sets how many decimal places to display the result. By default the command displays the first two decimal digits, but it also accepts: `0f`, `1f`, `2f`, `3f`, `4f`, `5f`, `6f` and just `f` (for all decimal digits).",
	Category:    "util",

	Arguments:    []string{"<number>f[optional]", "expression"},
	RequiredArgs: 1,

	DiscordPermissions: discordgo.PermissionSendMessages,

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
			msg.SendChannel("`%s` = %.0f", content, res)
		case "1f":
			msg.SendChannel("`%s` = %.1f", content, res)
		case "3f":
			msg.SendChannel("`%s` = %.3f", content, res)
		case "4f":
			msg.SendChannel("`%s` = %.4f", content, res)
		case "5f":
			msg.SendChannel("`%s` = %.5f", content, res)
		case "6f":
			msg.SendChannel("`%s` = %.6f", content, res)
		case "f":
			msg.SendChannel("`%s` = %f", content, res)
		default:
			msg.SendChannel("`%s` = %.2f", content, res)
		}
	},
}
