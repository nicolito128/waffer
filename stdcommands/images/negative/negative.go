package negative

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/superimage"
	"github.com/nicolito128/waffer/pkg/plugins"
	"github.com/nicolito128/waffer/pkg/plugins/supermessage"
)

var Command = plugins.Plugin[*discordgo.MessageCreate]{
	Name: "negative",
	Command: &plugins.CommandData{
		Description:  "Inverts the colors of an image.",
		Category:     "images",
		Arguments:    []string{"<url>.png/.jpg/.jpeg"},
		RequiredArgs: 1,
		Permissions: plugins.CommandPermissions{
			AllowDM: true,
			Require: discordgo.PermissionSendMessages & discordgo.PermissionAttachFiles,
		},
	},
	Handler: Handler,
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	sm := supermessage.New(s, m)

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

	ioutil.WriteFile(fmt.Sprintf("temp/negative.%s", neg.Format), buf.Bytes(), 0644)
	del := make(chan bool)

	go func() {
		sm.ChannelSendComplex(&discordgo.MessageSend{
			File: &discordgo.File{
				Name:        fmt.Sprintf("/temp/negative.%s", neg.Format),
				Reader:      bytes.NewReader(buf.Bytes()),
				ContentType: fmt.Sprintf("image/%s", neg.Format),
			},
		})

		del <- true
	}()

	if <-del {
		os.Remove(fmt.Sprintf("/temp/negative.%s", neg.Format))
	}
}
