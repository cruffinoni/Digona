package reaction

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/config"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"strings"
)

func editOriginalMessage(guildId, reactionId, roleId string) error {
	configuration := config.GetReactionMessageChannel(guildId)
	message, err := skeleton.Bot.GetSession().ChannelMessage(configuration.ChannelId, configuration.MessageId)
	if err != nil {
		skeleton.Bot.Error("unable to get discord message")
		return err
	}
	if reactMessages[message.ID] == nil {
		reactMessages[message.ID] = make(map[string]string)
	}
	if _, exists := reactMessages[message.ID][reactionId]; exists {
		delete(reactMessages[message.ID], reactionId)
		if err = skeleton.Bot.GetSession().MessageReactionsRemoveEmoji(configuration.ChannelId, message.ID, reactionId); err != nil {
			return err
		}
	} else {
		reactMessages[message.ID][reactionId] = roleId
		if err = skeleton.Bot.GetSession().MessageReactionAdd(configuration.ChannelId, message.ID, reactionId); err != nil {
			return err
		}
	}
	var description string
	for emoji, role := range reactMessages[configuration.MessageId] {
		if strings.Contains(emoji, ":") {
			description += fmt.Sprintf("<:%v>%v<@&%v>\n", emoji, delimiter, role)
		} else {
			description += fmt.Sprintf("%v%v<@&%v>\n", emoji, delimiter, role)
		}
	}
	_, err = skeleton.Bot.GetSession().ChannelMessageEditEmbed(configuration.ChannelId, configuration.MessageId, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       message.Embeds[0].Title,
		Description: description,
		Color:       message.Embeds[0].Color,
		Footer:      &discordgo.MessageEmbedFooter{Text: footerDescription},
	})
	return err
}

func AddRole(parser *parser.MessageParser) error {
	if len(parser.GetArguments()) != 2 {
		skeleton.Bot.SendDelayedMessage(parser.GetChannelId(), "Entrez l'emoji puis le rôle")
		return nil
	}
	var (
		err          error
		customEmojis []*discordgo.Emoji
		roles        []*discordgo.Role
		currentRole  *discordgo.Role
		args         = parser.GetArguments()
	)
	customEmojis, err = skeleton.Bot.GetSession().GuildEmojis(parser.GetGuildId())
	if err != nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Je ne peux pas récupérer les émojis personnalisés de ce serveur")
		return err
	}
	roles, err = skeleton.Bot.GetSession().GuildRoles(parser.GetGuildId())
	if err != nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}
	var currentEmojiId string
	if currentCustomEmoji := discord.FindEmojiFromRawEmojiId(customEmojis, args[0]); currentCustomEmoji == nil {
		currentEmojiId = args[0]
	} else {
		currentEmojiId = currentCustomEmoji.APIName()
	}
	if currentRole = discord.FindRoleFromRawRoleId(roles, args[1]); currentRole == nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), fmt.Sprintf("Impossible de trouver le rôle: '%v'"))
		return err
	}
	if err = editOriginalMessage(parser.GetGuildId(), currentEmojiId, currentRole.ID); err != nil {
		return err
	}
	return discord.DeleteMessage(parser.GetChannelId(), parser.GetDiscordMessage().ID)
}
