package invert

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
	Name:        "invertcolors",
	Aliases:     []string{"invc"},
	Description: "Inverts the colors of an image.",
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

		// Parsing the link
		_, format, err := superimage.ParseURL(link)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("**ParseURL error**: %s", err.Error()))
			return
		}

		m, _ := msg.SendChannelSafe("Processing image...")

		// Getting the image from URL
		img, err := superimage.GetImageFromURL(link)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("**GetImageFromURL error**: %s", err.Error()))
			return
		}

		// Creating a buffer to store the image
		// and then encoding it
		buf := new(bytes.Buffer)
		err = superimage.Encode(buf, img, format)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("**Encode error**: %s", err.Error()))
			return
		}

		// Writing the image to a file
		err = ioutil.WriteFile(fmt.Sprintf("./temp/tmp.%s", format), buf.Bytes(), 0666)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("**WriteFile error**: %s", err.Error()))
			return
		}

		// Inverting the colors
		inverseImg := superimage.InvertColors(img)

		// Creating a new buffer to store the image (required)
		// and then encoding it
		buf = new(bytes.Buffer)
		err = superimage.Encode(buf, inverseImg, format)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("**Encode inverse error**: %s", err.Error()))
			return
		}

		// Writing the inverse colors image to a file
		err = ioutil.WriteFile(fmt.Sprintf("./temp/tmp.%s", format), buf.Bytes(), 0666)
		if err != nil {
			msg.SendChannelSafe(fmt.Sprintf("**WriteFile inverse error**: %s", err.Error()))
			return
		}

		msg.SendMessageComplex(&discordgo.MessageSend{
			File: &discordgo.File{
				Name:        fmt.Sprintf("/temp/tmp.%s", format),
				Reader:      bytes.NewReader(buf.Bytes()),
				ContentType: fmt.Sprintf("image/%s", format),
			},
		})
	},
}
