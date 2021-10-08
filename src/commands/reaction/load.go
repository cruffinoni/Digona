package reaction

import (
	"github.com/cruffinoni/Digona/src/digona/config"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"strings"
)

func LoadReactionMessage(guildId string) error {
	reactionConfig := config.GetReactionMessageChannel(guildId)
	if reactionConfig.MessageId == "" {
		return nil
	}
	discMessage, err := skeleton.Bot.GetSession().ChannelMessage(reactionConfig.ChannelId, reactionConfig.MessageId)
	if err != nil {
		return err
	}
	lines := strings.Split(discMessage.Embeds[0].Description, "\n")
	reactMessages[reactionConfig.MessageId] = make(map[string]string)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineContent := strings.Split(line, delimiter)
		reactMessages[reactionConfig.MessageId][discord.GetReactionIdFromRawReactionId(lineContent[0])] = discord.GetRoleFromRawRoleId(lineContent[1])
	}
	return nil
}
