package hooks

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/user"
	"strings"
)

func isHiMessagePresent(words []string) bool {
	validHiMessage := []string{
		"hey", "hi", "salut", "hello", "wsh", "wesh", "bonjour", "bonsoir",
	}
	for _, w := range words {
		for _, vHi := range validHiMessage {
			if strings.ToLower(w) == vHi {
				return true
			}
		}
	}
	return false
}

func handleNormalMessage(parser *parser.MessageParser, message *discordgo.MessageCreate) {
	var err error
	if err = user.AddPoint(message.GuildID, message.Author.ID); err != nil {
		skeleton.Bot.Errorf("can't add point for new message: '%v'\n", err.Error())
		return
	}
	if isHiMessagePresent(parser.GetArguments()) {
		if err = skeleton.Bot.GetSession().MessageReactionAdd(message.ChannelID, message.ID, `👋`); err != nil {
			skeleton.Bot.Errorf("Can't add a reaction to the 'hi' message: %v", err)
		}
	}
	skeleton.Bot.Log("The author might not talk to me.\n")
	if user.IsPassingLevel(message.GuildID, message.Author.ID) {
		if err = user.SetNewLevel(message.GuildID, message.Author.ID); err != nil {
			skeleton.Bot.Errorf("can't set a new level to %v: %v\n", message.Author.Username, err)
		} else {
			skeleton.Bot.Logf("%v has level up to %v!\n", message.Author.Username, user.StoredUsers[message.GuildID][message.Author.ID].Level)
			_, err = skeleton.Bot.GetSession().ChannelMessageSend(message.ChannelID, fmt.Sprintf("%v tu viens de passer **au niveau %v** !", message.Author.Mention(), user.StoredUsers[message.GuildID][message.Author.ID].Level))
			if err != nil {
				skeleton.Bot.Errorf("can't send a message to %v: %v\n", message.Author.Username, err)
			}
		}
	}
}

func OnMessageCreated(_ *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == skeleton.Bot.GetID() {
		return
	}
	skeleton.Bot.Logf("> [%v] [%v]: %v\n", skeleton.Bot.GetGuildDataFromId(message.GuildID).Name, message.Author.Username, message.ContentWithMentionsReplaced())
	newParser := parser.New(message, skeleton.Bot.Logger)
	if !newParser.IsBotMentioned() {
		handleNormalMessage(newParser, message)
		return
	}
}
