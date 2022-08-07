package calc

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/go-calculator"
	"github.com/nicolito128/waffer/pkg/env"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/commands/flags"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Arguments = []*flags.Token{
	{
		Name:        "f",
		Short:       "f",
		Optional:    true,
		Description: "The number of decimal places to display the result.",
	},
	{
		Name:        "operation",
		Short:       "o",
		Default:     true,
		Description: "The operation to calculate.",
	},
}

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "calc",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:         "calc",
		Description:  fmt.Sprintf("A mathematical calculation command. Ex: `%scalc --op (10 / 2) + 1`, `%scalc --f 4 --op (1 / 9)`. The `--f` flag sets how many decimal places to display the result. By default the command displays the first two decimal digits, but it also accepts: `0`, `1`, `2`, `3`, `4`, `5` and `6`.", env.Get("BOT_PREFIX"), env.Get("BOT_PREFIX")),
		Category:     "util",
		Arguments:    strings.Split(flags.GetUsages(Arguments...), "\n"),
		RequiredArgs: 1,
		Permissions: &commands.CommandPermissions{
			Require: discordgo.PermissionSendMessages,
			AllowDM: true,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m.Message)
	f := flags.ParseString(sm.PlainContent(), Arguments...)

	if !flags.ResolveFlags(f, Arguments...) {
		sm.ChannelSend(fmt.Sprintf("**Invalid flags**.\n```\n%s\n```", flags.GetUsages(Arguments...)))
		return
	}

	var operation, decimals string
	operation = flags.GetFlag(f, "operation").Value

	if flags.HasFlag(f, "f") {
		decimals = flags.GetFlag(f, "f").Value
	} else {
		decimals = "2"
	}

	res, err := calculator.Resolve(operation)
	if err != nil {
		sm.ChannelSend("Error: %s", err)
		return
	}

	switch decimals {
	case "0":
		sm.ChannelSend("`%s` = %.0f", operation, res)
	case "1":
		sm.ChannelSend("`%s` = %.1f", operation, res)
	case "3":
		sm.ChannelSend("`%s` = %.3f", operation, res)
	case "4":
		sm.ChannelSend("`%s` = %.4f", operation, res)
	case "5":
		sm.ChannelSend("`%s` = %.5f", operation, res)
	case "6":
		sm.ChannelSend("`%s` = %.6f", operation, res)
	default:
		sm.ChannelSend("`%s` = %.2f", operation, res)
	}
}
