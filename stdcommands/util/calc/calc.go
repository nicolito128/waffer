package calc

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/go-calculator"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "calc",
	Command: &plugins.CommandData{
		Description:  "A mathematical calculation command. Use the example: `waf!calculator (10 / 2) + 1`, `waf!calculator 4f (1 / 9)`. The argument `<number>f` sets how many decimal places to display the result. By default the command displays the first two decimal digits, but it also accepts: `0f`, `1f`, `2f`, `3f`, `4f`, `5f`, `6f` and just `f` (for all decimal digits).",
		Category:     "util",
		Arguments:    []string{"<number>f[optional]", "expression"},
		RequiredArgs: 1,
		Permissions: plugins.CommandPermissions{
			Require: discordgo.PermissionSendMessages,
			AllowDM: true,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	args := sm.Arguments()

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
		sm.ChannelSend("Error: %s", err)
		return
	}

	switch decimals {
	case "0f":
		sm.ChannelSend("`%s` = %.0f", content, res)
	case "1f":
		sm.ChannelSend("`%s` = %.1f", content, res)
	case "3f":
		sm.ChannelSend("`%s` = %.3f", content, res)
	case "4f":
		sm.ChannelSend("`%s` = %.4f", content, res)
	case "5f":
		sm.ChannelSend("`%s` = %.5f", content, res)
	case "6f":
		sm.ChannelSend("`%s` = %.6f", content, res)
	case "f":
		sm.ChannelSend("`%s` = %f", content, res)
	default:
		sm.ChannelSend("`%s` = %.2f", content, res)
	}
}
