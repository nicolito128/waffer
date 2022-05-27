package holdingbooks

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
	animegirls "github.com/nicolito128/animegirls-holding-programming-books"
	"github.com/nicolito128/waffer/plugins/commands"
)

/**
 	* This command uses the images provided in the "Anime-Girls-Holding-Programming-Books"
 	* repository by cat-milk. All rights of the images to their respective authors.
	*
	* The command does not contain, or plan to contain, any +18 (adult) character images.
	*
	* Github repository: https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books
*/

var Command = &commands.WafferCommand{
	Name:        "girlholdingbook",
	Aliases:     []string{"ghb", "girlbook", "animebook", "girlprogrambook"},
	Description: "Random anime girl holding a programming book. Based repository: https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books",
	Category:    "fun",

	Arguments:    []string{"mention[optional]"},
	RequiredArgs: 0,

	DiscordPermissions: 0,

	RunInDM: false,

	RunFunc: func(dt *commands.HandlerData) {
		msg := dt.Message
		argument := strings.Join(msg.GetArguments(), " ")

		if argument == "" || argument == " " {
			rbIndex := rand.Intn(len(animegirls.Languages))
			rbLang := animegirls.Languages[rbIndex]

			im, err := animegirls.GetRandomImage(rbLang)
			if err != nil {
				msg.SendChannel("Error: ", err.Error())
			}

			msg.SendChannelEmbed(&discordgo.MessageEmbed{
				URL: im,
				Image: &discordgo.MessageEmbedImage{
					URL: im,
				},
			})
			return
		}

		im, err := animegirls.GetRandomImage(argument)
		if err != nil && im == "" {
			msg.SendChannel("No images found for that language.")
			msg.SendChannel(fmt.Sprintf("Try one of these languages: `%s`", strings.Join(animegirls.Languages, ", ")))
			return
		}

		msg.SendChannelEmbed(&discordgo.MessageEmbed{
			URL: im,
			Image: &discordgo.MessageEmbedImage{
				URL: im,
			},
		})
	},
}
