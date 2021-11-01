package opgg

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"net/url"
	"strings"
)

func buildUrl(pseudos []string) string {
	return "https://euw.op.gg/multi/query=" + url.QueryEscape(strings.Join(pseudos, ","))
}

func GetOPGGLink(message *parser.MessageParser) error {
	if len(message.GetArguments()) == 0 {
		skeleton.Bot.SendErrorMessage(message.GetChannelId(), "Cette commande s'utilise de deux manières différentes:\n"+
			"\t - Entrez les pseudonymes des utilisateurs en les séparants d'une virgule: \"Dayrion,Lindiana\"\n"+
			"\t - Entrez le message qu'apparaît quand vous entrez dans un lobby: \nDayrion joined the lobby\nLindiana joined the lobby")
		return nil
	}
	channelId := message.GetChannelId()
	lines := strings.Split(message.GetRawArguments(), "\n")
	var pseudos []string
	if len(lines) == 1 {
		pseudos = strings.Split(lines[0], ",")
		if len(pseudos) == 0 {
			skeleton.Bot.Errorf("no pseudo detected - content: %+#v\n", pseudos)
			skeleton.Bot.SendErrorMessage(channelId, "Une erreur s'est produite dans la détection des pseudos.")
			return nil
		}
	}
	pseudos = findAndDeleteCommonPattern(lines)
	if _, err := skeleton.Bot.GetSession().ChannelMessageSendEmbed(channelId, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "OP.GG",
		Description: fmt.Sprintf("Voici le lien OP.GG généré à partir des pseudos entrés:\n%v", buildUrl(pseudos)),
		Color:       skeleton.GenerateRandomMessageColor(),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://opgg-static.akamaized.net/icon/reverse.rectangle.png",
		},
	}); err != nil {
		skeleton.Bot.Errorf("Can't display a message into the channel '%v'. Error: '%v'\n", channelId, err)
	}
	return discord.DeleteMessage(channelId, message.GetDiscordMessage().ID)
}
