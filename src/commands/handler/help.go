package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

func HelpCommand(parser *parser.MessageParser) error {
	cmds := ""
	for key := range commandsListing {
		cmds += key + ", "
	}
	cmds = cmds[:len(cmds)-2]
	_, err := skeleton.Bot.GetSession().ChannelMessageSendEmbed(parser.GetChannelId(), &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Liste des commandes",
		Description: cmds,
		Color:       skeleton.GenerateRandomMessageColor(),
	})
	return err
}
