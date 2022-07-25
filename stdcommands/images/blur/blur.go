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
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

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
		Arguments:    []string{"<blur radio> <url>.png/.jpg/.jpeg"},
		RequiredArgs: 1,
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages & discordgo.PermissionAttachFiles,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)
	args := sm.Arguments()

	if len(args) < 1 {
		sm.ChannelSend("You must specify a link.")
		return
	}

	var link string
	var radio int

	if len(args) >= 2 {
		if args[0] == "" || args[0] == " " {
			radio = 2
		} else {
			val, err := strconv.Atoi(args[0])
			if err != nil {
				sm.ChannelSend("You must specify a valid blur radio. Error: ")
				return
			}

			radio = val
		}

		link = args[1]
	} else {
		link = strings.Join(args, "")
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
