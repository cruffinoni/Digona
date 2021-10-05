package config

import "github.com/cruffinoni/Digona/src/discord"

type ReactionConfig struct {
	ChannelId string `json:"channel_id"`
	MessageId string `json:"message_id"`
}

func UpdateReactionMessageId(guildId string, config ReactionConfig) {
	if configFiles[guildId].Configuration.Reaction.ChannelId != "" {
		if err := discord.DeleteMessage(configFiles[guildId].Configuration.Reaction.ChannelId, configFiles[guildId].Configuration.Reaction.MessageId); err != nil {
			log.Errorf("can't delete the previous reaction message (guild %v)\n", guildId)
		}
	}
	configFiles[guildId].Configuration.Reaction = config
	if err := Save(guildId); err != nil {
		log.Errorf("unable to save config file (%v) => %v\n", guildId, err)
	}
}

func GetReactionMessageChannel(guildId string) ReactionConfig {
	return configFiles[guildId].Configuration.Reaction
}
