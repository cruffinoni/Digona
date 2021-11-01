package opgg

import (
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"net/url"
	"strings"
)

func buildUrl(pseudos []string) string {
	return "https://euw.op.gg/multi/query=" + url.QueryEscape(strings.Join(pseudos, ","))
}

func GetOPGGLink(message *parser.MessageParser) error {
	lines := strings.Split(message.GetRawArguments(), "\n")
	pseudos := make([]string, 0, 5)
	for _, l := range lines {
		if len(l) == 0 {
			skeleton.Bot.Logf("Empty line? %v\n", l)
			continue
		}
		words := strings.Split(l, " ")
		skeleton.Bot.Logf("Words? %v\n", words)
		if len(words) == 0 {
			continue
		}
		pseudos = append(pseudos, words[0])
	}
	skeleton.Bot.Logf("Pseudos found: %+v\n", pseudos)
	skeleton.Bot.Logf("URL\n%v\n", buildUrl(pseudos))
	return nil
}
