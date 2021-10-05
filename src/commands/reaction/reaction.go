package reaction

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/config"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"regexp"
	"strings"
)

var (
	reactMessages = make(map[string]map[string]string)
)

const (
	delimiter = "⟹"
)

func GetRoleFromMessageReaction(messageId, reactionId string) string {
	if reactMessages[messageId] != nil && reactMessages[messageId][reactionId] != "" {
		return reactMessages[messageId][reactionId]
	}
	return ""
}

func formatMessage(args []string) string {
	var message string
	for i, content := range args {
		message += content
		if i+1 != len(args) {
			message += " "
		}
	}
	return message
}

func setupMessageAndReactions(parser *parser.MessageParser, messageContent string, reactions map[string]string) error {
	var description string
	for emoji, role := range reactions {
		if strings.Contains(emoji, ":") {
			description += fmt.Sprintf("<:%v>%v<@&%v>\n", emoji, delimiter, role)
		} else {
			description += fmt.Sprintf("%v%v<@&%v>\n", emoji, delimiter, role)
		}
	}
	message, err := skeleton.Bot.GetSession().ChannelMessageSendEmbed(parser.GetChannelId(), &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       messageContent,
		Description: description,
		Color:       skeleton.GenerateRandomMessageColor(),
	})
	reactMessages[message.ID] = reactions
	if err != nil {
		delete(reactMessages, message.ID)
		return err
	}
	for emoji := range reactions {
		if err = skeleton.Bot.GetSession().MessageReactionAdd(parser.GetChannelId(), message.ID, emoji); err != nil {
			secondErr := skeleton.Bot.GetSession().ChannelMessageDelete(parser.GetChannelId(), message.ID)
			if secondErr != nil {
				skeleton.Bot.Errorf("Cannot delete message id '%v': %v\n", message.ID, secondErr)
			}
			skeleton.Bot.SendMessage(parser.GetChannelId(), "Une erreur est survenue, réessayez plus tard")
			return err
		}
	}
	config.UpdateReactionMessageId(parser.GetGuildId(), config.ReactionConfig{
		ChannelId: parser.GetChannelId(),
		MessageId: message.ID,
	})
	return nil
}

func Role(parser *parser.MessageParser) error {
	if len(parser.GetArguments()) < 3 {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Format de la command: [CMD] [ROLE] [REACTION] [MESSAGE]")
		return nil
	}
	customEmojis, err := skeleton.Bot.GetSession().GuildEmojis(parser.GetGuildId())
	if err != nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}
	roles, err := skeleton.Bot.GetSession().GuildRoles(parser.GetGuildId())
	if err != nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}

	var currentRole *discordgo.Role = nil
	var currentEmojiId string
	var messageContent string
	listedReactions := make(map[string]string)

	for k, i := range parser.GetArguments() {
		if currentRole == nil {
			if matched, err := regexp.Match("<@&\\d{18}>", []byte(i)); err != nil {
				skeleton.Bot.SendMessage(parser.GetChannelId(), "Une erreur s'est produite, réessayez plus tard")
				return err
			} else if !matched {
				messageContent = formatMessage(parser.GetArguments()[k:])
				break
			}
			if currentRole = discord.FindRoleFromRawRoleId(roles, i); currentRole == nil {
				skeleton.Bot.SendMessage(parser.GetChannelId(), "Impossible de trouver le rôle: '"+i+"'")
				return err
			}
		} else {
			if currentCustomEmoji := discord.FindEmojiFromRawEmojiId(customEmojis, i); currentCustomEmoji == nil {
				currentEmojiId = i
			} else {
				currentEmojiId = currentCustomEmoji.APIName()
			}
			listedReactions[currentEmojiId] = currentRole.ID
			currentRole = nil
			currentEmojiId = ""
		}
	}
	if currentRole != nil || currentEmojiId != "" {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Le nombre d'argument est incorrect. Est-ce que le message de fin manque t-il?")
		return nil
	}
	return setupMessageAndReactions(parser, messageContent, listedReactions)
}
