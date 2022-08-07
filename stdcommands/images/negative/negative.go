package negative

import (
	"bytes"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/superimage"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/commands"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

// Create a commands.WafferCommand
var Command = &commands.WafferCommand{
	Plugin: &plugins.Plugin[*discordgo.MessageCreate]{
		Name:    "negative",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:         "negative",
		Description:  "Inverts the colors of an image.",
		Category:     "images",
		Arguments:    []string{"<url>.png/.jpg/.jpeg"},
		RequiredArgs: 1,
		Permissions: &commands.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages & discordgo.PermissionAttachFiles,
		},
	},
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m.Message)

	link := sm.PlainContent()

	if link == "" || link == " " {
		sm.ChannelSend("You need to specify a link.")
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

	neg := superimage.Negative(img)
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, neg, nil)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Encoding error: %s", err.Error()))
		return
	}

	direction := fmt.Sprintf("temp/negative.%s", neg.Format)
	file, err := os.OpenFile(direction, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Creating file error: %s", err.Error()))
		return
	}
	defer file.Close()

	send := make(chan bool)
	del := make(chan bool)

	go func() {
		_, err = file.Write(buf.Bytes())
		if err != nil {
			sm.ChannelSend(fmt.Sprintf("Writing file error: %s", err.Error()))
			send <- false
			return
		}

		send <- true
	}()

	go func() {
		sm.ChannelSend("Processing image...")

		if <-send {
			_, err := sm.ChannelSendComplex(&discordgo.MessageSend{
				File: &discordgo.File{
					Name:        fmt.Sprintf("/temp/negative.%s", neg.Format),
					Reader:      bytes.NewReader(buf.Bytes()),
					ContentType: fmt.Sprintf("image/%s", neg.Format),
				},
			})

			if err != nil {
				sm.ChannelSend("Error sending image.")
				del <- false
				return
			}

		} else {
			del <- false
			return
		}

		del <- true
	}()

	if <-del {
		os.Remove(direction)
	}

	close(send)
	close(del)
}
