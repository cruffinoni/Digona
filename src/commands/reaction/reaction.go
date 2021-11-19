package reaction

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/config"
	"github.com/cruffinoni/Digona/src/database/models"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
)

var (
	reactMessages     = make(map[string]map[string]string)
	footerDescription = "Choisissez un rôle en interagissant avec les réactions. Utilisez " +
		"'react-add' pour ajouter des rôles."
)

const (
	delimiter = " ⟹ "
)

func GetRoleFromMessageReaction(messageId, reactionId string) string {
	if reactMessages[messageId] != nil && reactMessages[messageId][reactionId] != "" {
		return reactMessages[messageId][reactionId]
	}
	return ""
}

func setupMessageAndReactions(channelId, guildId, messageContent string) error {
	message, err := skeleton.Bot.GetSession().ChannelMessageSendEmbed(channelId, &discordgo.MessageEmbed{
		Type:   discordgo.EmbedTypeRich,
		Footer: &discordgo.MessageEmbedFooter{Text: footerDescription},
		Title:  messageContent,
		Color:  skeleton.GenerateRandomMessageColor(),
	})
	reactMessages[message.ID] = make(map[string]string)
	config.UpdateReactionMessageId(guildId, models.ReactionConfig{
		ChannelId: channelId,
		MessageId: message.ID,
	})
	return err
}

func Role(parser *parser.MessageParser) error {
	if len(parser.GetArguments()) < 1 {
		skeleton.Bot.SendMessageWithNoTitle(parser.GetChannelId(), "Entrez le message à afficher.")
		return nil
	}

	var message string
	for i, content := range parser.GetArguments() {
		message += content
		if i+1 != len(parser.GetArguments()) {
			message += " "
		}
	}
	if err := setupMessageAndReactions(parser.GetChannelId(), parser.GetGuildId(), message); err != nil {
		return err
	}
	skeleton.Bot.SendDelayedMessage(parser.GetChannelId(), "Vous pouvez ajouter des rôles avec la commandes: 'react-add'!")
	return discord.DeleteMessage(parser.GetChannelId(), parser.GetDiscordMessage().ID)
}
