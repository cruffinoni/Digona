package commands

import (
	"errors"
	"fmt"
	"github.com/Digona/src/digona"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

const (
	minAmountDeleteMsg = 1
	maxAmountDeleteMsg = 100
)

func deleteUserMessages(parser *MessageParser, channelMessage []*discordgo.Message, maxMessage int) error {
	var deletedMsg int
	for count, message := range channelMessage {
		if message.Author.Username == parser.message.Author.Username {
			err := digona.Bot.GetSession().ChannelMessageDelete(parser.channel, message.ID)
			if err != nil {
				_ = digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je n'ai réussi à supprimer que %v de tes messages %v.", count, parser.message.Author.Mention()))
				return err
			}
			deletedMsg++
			if deletedMsg >= maxMessage {
				return nil
			}
		}
	}
	return nil
}

func deleteLastMessages(parser *MessageParser, channelMessage []*discordgo.Message) error {
	for count, eachMessage := range channelMessage {
		err := digona.Bot.GetSession().ChannelMessageDelete(parser.channel, eachMessage.ID)
		if err != nil {
			_ = digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je n'ai réussi à supprimer que %v message(s).", count))
			return err
		}
	}
	return nil
}

func getNumberOfMessageToDelete(message []string) (int, error) {
	for _, arg := range message {
		if number, err := strconv.Atoi(arg); err == nil {
			return number, nil
		}
	}
	return 0, errors.New("no valid arg found")
}

func redirectDelete(parser *MessageParser) error {
	if len(parser.args) == 0 {
		return digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je n'ai pas compris ce que tu veux dire avec la commande %v!", parser.Command))
	}
	count, err := getNumberOfMessageToDelete(parser.args)
	if count == 0 && err != nil {
		return digona.Bot.DisplayError(parser.channel, "Tu dois entrer le nombre de message que tu souhaites supprimer.")
	}
	if count < minAmountDeleteMsg || count > (maxAmountDeleteMsg - 1) {
		return digona.Bot.DisplayError(parser.channel, fmt.Sprintf("Je ne peux supprimer qu'entre %v et %v messages à la fois.", minAmountDeleteMsg, maxAmountDeleteMsg - 1))
	}
	fmt.Printf("So: count: %v & tagged? %v\n", count, parser.IsTaggingHimself())
	if parser.IsTaggingHimself() {
		allMessages, err := digona.Bot.GetSession().ChannelMessages(parser.channel, maxAmountDeleteMsg, "", "", parser.message.ID)
		if err != nil {
			_ = digona.Bot.DisplayError(parser.channel, "Je n'arrive pas à supprimer tes messages... Réessayes dans quelques minutes " + parser.message.Author.Mention() + ".")
			return err
		}
		return deleteUserMessages(parser, allMessages, count + 1)
	}
	allMessages, err := digona.Bot.GetSession().ChannelMessages(parser.channel, count + 1, "", "", parser.message.ID)
	if err != nil {
		_ = digona.Bot.DisplayError(parser.channel, "Je n'arrive pas à supprimer les messages... Réessayes dans quelques minutes." + parser.message.Author.Mention() + ".")
		return err
	}
	return deleteLastMessages(parser, allMessages)
}