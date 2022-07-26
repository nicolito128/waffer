package blur

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/superimage"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/commands/flags"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Arguments = []*flags.Token{
	{
		Name:        "radio",
		Short:       "r",
		Optional:    true,
		Description: "The blur radio.",
	},
	{
		Name:        "url",
		Short:       "u",
		Default:     true,
		Description: "The image URL.",
	},
}

var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "blur",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:         "blur",
		Description:  "Blur an image.",
		Category:     "images",
		Arguments:    strings.Split(flags.GetUsages(Arguments...), "\n"),
		RequiredArgs: 1,
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages & discordgo.PermissionAttachFiles,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	args := flags.ParseString(sm.PlainContent(), Arguments...)

	if !flags.ResolveFlags(args, Arguments...) {
		sm.ChannelSend(fmt.Sprintf("**Invalid flags**.\n```\n%s\n```", flags.GetUsages(Arguments...)))
		return
	}

	var r, link string
	if flags.HasFlag(args, "url") {
		link = flags.GetFlag(args, "url").Value
	} else {
		if len(sm.Create.Attachments) >= 1 {
			link = sm.Create.Attachments[0].URL
		} else {
			sm.ChannelSend("**Missing image URL or attachment.**")
			return
		}
	}

	var radio int
	if flags.HasFlag(args, "radio") {
		r = flags.GetFlag(args, "radio").Value

		result, err := strconv.Atoi(r)
		if err != nil {
			sm.ChannelSend("**Error**. You must specify a valid blur radio")
			return
		}

		radio = result
	} else {
		radio = 2
	}

	if link == "" || link == " " {
		sm.ChannelSend("You must to specify a valid link.")
		return
	}

	img, err := superimage.GetByURL(link)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Gettin URL error: %s", err.Error()))
		return
	}

	buf := new(bytes.Buffer)
	err = superimage.Encode(buf, img, nil)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Encoding error: %s", err.Error()))
		return
	}

	blurred, err := superimage.Blur(img, radio)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Blur error: %s", err.Error()))
		return
	}

	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, blurred, nil)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Encoding error: %s", err.Error()))
		return
	}

	ioutil.WriteFile(fmt.Sprintf("/temp/blurred.%s", blurred.Format), buf.Bytes(), 0644)
	del := make(chan bool)

	go func() {
		sm.ChannelSend("Processing image...")
		_, err := sm.ChannelSendComplex(&discordgo.MessageSend{
			File: &discordgo.File{
				Name:        fmt.Sprintf("/temp/blurred.%s", blurred.Format),
				Reader:      bytes.NewReader(buf.Bytes()),
				ContentType: fmt.Sprintf("image/%s", blurred.Format),
			},
		})

		if err != nil {
			sm.ChannelSend("Error sending image.")
		}

		del <- true
	}()

	if <-del {
		os.Remove(fmt.Sprintf("/temp/blurred.%s", blurred.Format))
	}
}
