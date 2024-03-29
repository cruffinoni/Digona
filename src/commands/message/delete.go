package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/discord"
)

const (
	minAmountDeleteMsg = 1
	maxAmountDeleteMsg = 100
)

/*
// can be useful later

func deleteUserMessages(parser *parser.MessageParser, channelMessage []*discordgo.Message, maxMessage int) error {
	var deletedMsg int
	for count, message := range channelMessage {
		if message.Author.Username == parser.GetDiscordMessage().Author.Username {
			err := skeleton.Bot.GetSession().ChannelMessageDelete(parser.GetChannelId(), message.ID)
			if err != nil {
				skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), fmt.Sprintf("Je n'ai réussi à supprimer que %v de tes messages %v.", count, parser.GetDiscordMessage().Author.Mention()))
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
*/

func deleteLastMessages(channelId string, channelMessage []*discordgo.Message) (uint, error) {
	var count uint
	for _, eachMessage := range channelMessage {
		err := discord.DeleteMessage(channelId, eachMessage.ID)
		if err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

//func getNumberOfMessageToDelete(message []string) (int, error) {
//	for _, arg := range message {
//		if number, err := strconv.Atoi(arg); err == nil {
//			return number, nil
//		}
//	}
//	return 0, errors.New("no valid arg found")
//}

//func RedirectDelete(parser *parser.MessageParser) error {
//	if len(parser.GetArguments()) == 0 {
//		skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), "Tu dois entrer le nombre de message à supprimer!")
//		return nil
//	}
//	count, err := getNumberOfMessageToDelete(parser.GetArguments())
//	if count == 0 && err != nil {
//		skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), "Tu dois entrer le nombre de message que tu souhaites supprimer.")
//		return err
//	}
//	if count < minAmountDeleteMsg || count > (maxAmountDeleteMsg-1) {
//		skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), fmt.Sprintf("Je ne peux supprimer qu'entre %v et %v messages à la fois.", minAmountDeleteMsg, maxAmountDeleteMsg-1))
//		return nil
//	}
//	if parser.IsTaggingHimself() {
//		allMessages, err := skeleton.Bot.GetSession().ChannelMessages(parser.GetChannelId(), maxAmountDeleteMsg, "", "", parser.GetDiscordMessage().ID)
//		if err != nil {
//			skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), "Je n'arrive pas à supprimer tes messages... Essayez dans quelques minutes "+parser.GetDiscordMessage().Author.Mention()+".")
//			return err
//		}
//		return deleteUserMessages(parser, allMessages, count+1)
//	}
//	allMessages, err := skeleton.Bot.GetSession().ChannelMessages(parser.GetChannelId(), count+1, "", "", parser.GetDiscordMessage().ID)
//	if err != nil {
//		skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), "Je n'arrive pas à supprimer les messages... Essayez dans quelques minutes."+parser.GetDiscordMessage().Author.Mention()+".")
//		return err
//	}
//	return deleteLastMessages(parser, allMessages)
//}
