package reflect

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		Name:    "reflect",
		Type:    plugins.MessageCreateType,
		Handler: Handler,
	},

	Data: &commands.CommandData{
		Name:         "reflect",
		Description:  "Reflect an image.",
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

	reflected := superimage.Reflect(img)
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, reflected, nil)
	if err != nil {
		sm.ChannelSend(fmt.Sprintf("Encoding error: %s", err.Error()))
		return
	}

	ioutil.WriteFile(fmt.Sprintf("/temp/reflected.%s", reflected.Format), buf.Bytes(), 0644)
	del := make(chan bool)

	go func() {
		sm.ChannelSend("Processing image...")
		_, err := sm.ChannelSendComplex(&discordgo.MessageSend{
			File: &discordgo.File{
				Name:        fmt.Sprintf("/temp/reflected.%s", reflected.Format),
				Reader:      bytes.NewReader(buf.Bytes()),
				ContentType: fmt.Sprintf("image/%s", reflected.Format),
			},
		})

		if err != nil {
			sm.ChannelSend("Error sending image.")
		}

		del <- true
	}()

	if <-del {
		os.Remove(fmt.Sprintf("/temp/reflected.%s", reflected.Format))
	}
}
