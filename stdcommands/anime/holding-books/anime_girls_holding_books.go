package holdingbooks

/**
 	* This command uses the images provided in the "Anime-Girls-Holding-Programming-Books"
 	* repository by cat-milk. All rights of the images to their respective authors.
	*
	* The command does not contain, or plan to contain, any +18 (adult) character images.
	*
	* Github repository: https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books
*/

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/nicolito128/waffer/plugins/commands"
)

type DataJSON struct {
	AI        []string `json:"ai"`
	Go        []string `json:"go"`
	Ocaml     []string `json:"ocaml"`
	Haskell   []string `json:"haskell"`
	Smalltalk []string `json:"smalltalk"`
	Elixir    []string `json:"elixir"`
	CSharp    []string `json:"csharp"`
}

var languages = []string{"ai", "go", "ocaml", "haskell", "smalltalk", "elixir", "csharp"}
var aghpb DataJSON

var Command = &commands.WafferCommand{
	Name:        "girlholdingbook",
	Aliases:     []string{"ghb", "girlbook", "animebook"},
	Description: "Random anime girl holding a programming book.",
	Category:    "fun",

	Arguments:    []string{"mention[optional]"},
	RequiredArgs: 0,

	DiscordPermissions: 0,

	RunInDM: false,

	RunFunc: func(dt *commands.HandlerData) {
		msg := dt.Message

		// Anime girls holding programming book
		fileData, _ := os.ReadFile("./stdcommands/anime/holding-books/data.json")
		// Parsing json data into aghpb
		json.Unmarshal(fileData, &aghpb)

		argument := strings.Join(msg.GetArguments(), " ")

		if argument == "" || argument == " " {
			rbIndex := rand.Intn(len(languages))
			link := getRandomLinkByLang(languages[rbIndex])
			msg.SendChannel(link)
		} else {
			link := getRandomLinkByLang(strings.ToLower(argument))
			if link == "" {
				msg.SendChannel(fmt.Sprintf("Languages not available. Try one of this: `%s`", strings.Join(languages, " | ")))
			} else {
				msg.SendChannel(link)
			}
		}
	},
}

func getRandomLinkByLang(lang string) string {
	switch lang {
	case "ai":
		rbIndex := rand.Intn(len(aghpb.AI))
		return aghpb.AI[rbIndex]
	case "go":
		rbIndex := rand.Intn(len(aghpb.Go))
		return aghpb.Go[rbIndex]
	case "ocaml":
		rbIndex := rand.Intn(len(aghpb.Ocaml))
		return aghpb.Ocaml[rbIndex]
	case "haskell":
		rbIndex := rand.Intn(len(aghpb.Haskell))
		return aghpb.Haskell[rbIndex]
	case "smalltalk":
		rbIndex := rand.Intn(len(aghpb.Smalltalk))
		return aghpb.Smalltalk[rbIndex]
	case "elixir":
		rbIndex := rand.Intn(len(aghpb.Elixir))
		return aghpb.Elixir[rbIndex]
	case "csharp", "c#":
		rbIndex := rand.Intn(len(aghpb.CSharp))
		return aghpb.CSharp[rbIndex]
	default:
		return ""
	}
}
