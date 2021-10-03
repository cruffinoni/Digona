package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"regexp"
	"time"
)

var (
	reactMessages = make(map[string]map[string]string)
)

func GetRoleFromMessageReaction(messageId, reactionId string) string {
	if reactMessages[messageId] != nil && reactMessages[messageId][reactionId] != "" {
		return reactMessages[messageId][reactionId]
	}
	return ""
}

func retrieveRole(roles []*discordgo.Role, reference string) *discordgo.Role {
	reference = reference[3 : len(reference)-1]
	for _, r := range roles {
		if r.ID == reference {
			return r
		}
	}
	return nil
}

func retrieveEmoji(customEmojis []*discordgo.Emoji, reference string) *discordgo.Emoji {
	reference = reference[1 : len(reference)-1]
	for _, r := range customEmojis {
		fullRef := ":" + r.Name + ":" + r.ID
		if fullRef == reference {
			return r
		}
	}
	return nil
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

func setupMessageAndReactions(parser *MessageParser, messageContent string, reactions map[string]string) error {
	message, err := skeleton.Bot.GetSession().ChannelMessageSendEmbed(parser.channel, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Réagissez à ce message",
		Description: messageContent,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       skeleton.GenerateRandomMessageColor(),
	})
	reactMessages[message.ID] = reactions
	if err != nil {
		delete(reactMessages, message.ID)
		return err
	}
	for emoji := range reactions {
		if err = skeleton.Bot.GetSession().MessageReactionAdd(parser.channel, message.ID, emoji); err != nil {
			secondErr := skeleton.Bot.GetSession().ChannelMessageDelete(parser.channel, message.ID)
			if secondErr != nil {
				skeleton.Bot.Errorf("Cannot delete message id '%v': %v\n", message.ID, secondErr)
			}
			skeleton.Bot.SendMessage(parser.channel, "Une erreur est survenue, réessayez plus tard")
			return err
		}
	}
	return nil
}

func Role(parser *MessageParser) error {
	if len(parser.args) < 3 {
		skeleton.Bot.SendMessage(parser.channel, "Format de la command: [CMD] [ROLE] [REACTION] [MESSAGE]")
		return nil
	}
	customEmojis, err := skeleton.Bot.GetSession().GuildEmojis(parser.guildId)
	if err != nil {
		skeleton.Bot.SendMessage(parser.channel, "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}
	roles, err := skeleton.Bot.GetSession().GuildRoles(parser.guildId)
	if err != nil {
		skeleton.Bot.SendMessage(parser.channel, "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}

	var currentRole *discordgo.Role = nil
	var currentEmojiId string
	var messageContent string
	listedReactions := make(map[string]string)

	for k, i := range parser.args {
		if currentRole == nil {
			if matched, err := regexp.Match("<@&\\d{18}>", []byte(i)); err != nil {
				skeleton.Bot.SendMessage(parser.channel, "Une erreur s'est produite, réessayez plus tard")
				return err
			} else if !matched {
				messageContent = formatMessage(parser.args[k:])
				break
			}
			if currentRole = retrieveRole(roles, i); currentRole == nil {
				skeleton.Bot.SendMessage(parser.channel, "Impossible de trouver le rôle: '"+i+"'")
				return err
			}
		} else {
			if currentCustomEmoji := retrieveEmoji(customEmojis, i); currentCustomEmoji == nil {
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
		skeleton.Bot.SendMessage(parser.channel, "Le nombre d'argument est incorrect. Est-ce que le message de fin manque t-il?")
		return nil
	}
	return setupMessageAndReactions(parser, messageContent, listedReactions)
}
