package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/logger"
	"strings"
)

type CommandHandler func(*MessageParser) error

type MessageParser struct {
	command     string
	author      *discordgo.User
	channel     string
	args        []string
	Handler     CommandHandler
	isMentioned bool
	message     *discordgo.Message
	guildId     string
	logger      logger.Logger
}

func checkIsBotMentioned(tab []*discordgo.User) bool {
	for _, user := range tab {
		if user.ID == skeleton.Bot.GetID() {
			return true
		}
	}
	return false
}

var commandsListing = map[string]CommandHandler{
	"delete":       redirectDelete,
	"react":        Role,
	"qr-code":      GenerateQrCode,
	"default-role": SetDefaultRole,
}

func New(message *discordgo.MessageCreate, logger logger.Logger) (parser *MessageParser) {
	parser = &MessageParser{
		message:     message.Message,
		channel:     message.ChannelID,
		author:      message.Author,
		isMentioned: checkIsBotMentioned(message.Mentions),
		logger:      logger,
		guildId:     message.GuildID,
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
				if j != i && realContent != skeleton.Bot.GetMention() && realContent != "" {
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
