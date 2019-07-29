package commands

import (
	"fmt"
	"github.com/Digona/src/digona"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type MessageParser struct {
	Command     string
	channel     string
	args        []string
	handler     commandHandler
	isMentioned bool
	message     *discordgo.Message
}

func isBotMentionned(tab []*discordgo.User) bool {
	for _, user := range tab {
		if user.ID == digona.Bot.GetID() {
			return true
		}
	}
	return false
}

func New(message *discordgo.MessageCreate) (parser *MessageParser) {
	parser = &MessageParser {
		message:	message.Message,
		channel:	message.ChannelID,
	}
	if !isBotMentionned(message.Mentions) {
		return
	}
	parser.isMentioned = true
	msgContent := strings.Split(message.Content, " ")
	fmt.Printf("Args splitted: %+v & %v\n", msgContent, msgContent[0])
	for i, word := range msgContent {
		if function, exists := userCommands[word]; exists {
			parser.handler = function
			parser.Command = word
			parser.args = msgContent[i + 1:]
			return
		}
	}
	return
}

func (parser *MessageParser) IsTaggingHimself() bool {
	for _, arg := range parser.args {
		if arg == "@me" || arg == parser.message.Author.Mention() {
			return true
		}
	}
	return false
}