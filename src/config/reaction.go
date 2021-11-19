package config

import (
	"github.com/cruffinoni/Digona/src/database/models"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
)

func UpdateReactionMessageId(guildId string, config models.ReactionConfig) {
	if configFiles[guildId].Configuration.Reaction.ChannelId != "" {
		if err := discord.DeleteMessage(configFiles[guildId].Configuration.Reaction.ChannelId, configFiles[guildId].Configuration.Reaction.MessageId); err != nil {
			log.Errorf("can't delete the previous reaction message (guild %v)\n", guildId)
		}
	}
	configFiles[guildId].Configuration.Reaction = config
	if err := skeleton.Bot.GetDatabase().SaveConfigFile(guildId, configFiles[guildId]); err != nil {
		log.Errorf("unable to save config file (%v) => %v\n", guildId, err)
	}
}

func GetReactionMessageChannel(guildId string) models.ReactionConfig {
	return configFiles[guildId].Configuration.Reaction
}
