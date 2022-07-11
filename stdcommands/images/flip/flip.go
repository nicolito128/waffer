package flip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicolito128/superimage"
	"github.com/nicolito128/waffer/plugins/commands"
)

var Command = &commands.WafferCommand{
	Name:        "flip",
	Aliases:     []string{},
	Description: "Inverts an image vertically.",
	Category:    "images",

	Arguments:    []string{"URL"},
	RequiredArgs: 1,

	OwnerOnly:          false,
	DiscordPermissions: discordgo.PermissionSendMessages,

	RunInDM: true,

	RunFunc: func(data *commands.HandlerData) {
		msg := data.Message
		link := strings.Trim(strings.Join(msg.GetArguments(), " "), " ")

		if link == "" || link == " " {
			msg.SendChannelSafe("You need to specify a link.")
			return
		}

		img, err := superimage.GetByURL(link)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("Gettin URL error: %s", err.Error()))
			return
		}

		buf := new(bytes.Buffer)
		err = superimage.Encode(buf, img, nil)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("Encoding error: %s", err.Error()))
			return
		}

		flipped := superimage.Flip(img)
		buf = new(bytes.Buffer)
		err = superimage.Encode(buf, flipped, nil)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("Encoding error: %s", err.Error()))
			return
		}

		ioutil.WriteFile("temp/flip.png", buf.Bytes(), 0644)

		msg.SendMessageComplex(&discordgo.MessageSend{
			File: &discordgo.File{
				Name:        fmt.Sprintf("/temp/flip.%s", flipped.Format),
				Reader:      bytes.NewReader(buf.Bytes()),
				ContentType: fmt.Sprintf("image/%s", flipped.Format),
			},
		})
	},
}
