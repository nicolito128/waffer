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

	OwnerOnly:          false,
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
			msg.SendChannelSafe("Error: %s", err)
			return
		}

		switch decimals {
		case "0f":
			msg.SendChannelSafe("`%s` = %.0f", content, res)
		case "1f":
			msg.SendChannelSafe("`%s` = %.1f", content, res)
		case "3f":
			msg.SendChannelSafe("`%s` = %.3f", content, res)
		case "4f":
			msg.SendChannelSafe("`%s` = %.4f", content, res)
		case "5f":
			msg.SendChannelSafe("`%s` = %.5f", content, res)
		case "6f":
			msg.SendChannelSafe("`%s` = %.6f", content, res)
		case "f":
			msg.SendChannelSafe("`%s` = %f", content, res)
		default:
			msg.SendChannelSafe("`%s` = %.2f", content, res)
		}
	},
}
