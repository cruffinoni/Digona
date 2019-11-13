package commands

import (
	"github.com/Digona/digona"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandHandler func(*MessageParser) error

type MessageParser struct {
	command     string
	channel     string
	args        []string
	Handler     CommandHandler
	isMentioned bool
	message     *discordgo.Message
}

var commandsListing = map[string]CommandHandler{
	"delete":  RedirectDelete,
	"players": ShowMostPlayedChamp,
}

func checkIsBotMentioned(tab []*discordgo.User) bool {
	for _, user := range tab {
		if user.ID == digona.Bot.GetID() {
			return true
		}
	}
	return false
}

func New(message *discordgo.MessageCreate) (parser *MessageParser) {
	parser = &MessageParser{
		message:     message.Message,
		channel:     message.ChannelID,
		isMentioned: checkIsBotMentioned(message.Mentions),
	}
	if !checkIsBotMentioned(message.Mentions) {
		return
	}
	parser.isMentioned = true
	msgContent := strings.Split(message.Content, " ")
	for i, word := range msgContent {
		if function, exists := commandsListing[strings.ToLower(word)]; exists {
			parser.Handler = function
			parser.command = word
			for j, realContent := range msgContent {
				if j != i && realContent != digona.Bot.GetMention() {
					parser.args = append(parser.args, realContent)
				}
			}
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

func (parser *MessageParser) GetArguments() []string {
	return parser.args
}

func (parser *MessageParser) GetChannelId() string {
	return parser.channel
}

func (parser *MessageParser) GetDiscordMessage() *discordgo.Message {
	return parser.message
}

func (parser *MessageParser) IsBotMentioned() bool {
	return parser.isMentioned
}

func (parser *MessageParser) GetOriginalCommand() string {
	return parser.command
}
