package commands

import (
	"fmt"
	"github.com/Digona/src/digona"
	"strconv"
)

const (
	minAmountDeleteMsg = 1
	maxAmountDeleteMsg = 100
)

func deleteLastMessages(parser *MessageParser) error {
	if len(parser.Args) == 0 {
		return digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je n'ai pas compris ce que tu veux dire avec la commande %v!", parser.Command))
	}
	count, err := strconv.Atoi(parser.Args[0])
	if err != nil {
		return digona.Bot.DisplayError(parser.channel, "Assure d'entrer un nombre après le nom de la commande.")
	}
	if count < minAmountDeleteMsg || count > (maxAmountDeleteMsg - 1) {
		return digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je ne peux supprimer qu'entre %v et %v messages à la fois.", minAmountDeleteMsg, maxAmountDeleteMsg - 1))
	}
	allMessages, err := digona.Bot.Session.ChannelMessages(parser.channel, count, "", "", parser.message.ID)
	if err != nil {
		_ = digona.Bot.DisplayError(parser.channel, "Je n'arrive pas à supprimer les messages... Réessayes dans quelques minutes.")
		return err
	}
	for count, eachMessage := range allMessages {
		err := digona.Bot.Session.ChannelMessageDelete(parser.channel, eachMessage.ID)
		if err != nil {
			_ = digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je n'ai réussi à supprimer que %v message(s).", count))
			return err
		}
	}
	return nil
}