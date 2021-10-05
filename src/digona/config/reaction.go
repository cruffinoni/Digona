package config

func UpdateReactionMessageId(guildId, messageId string) {
	configFiles[guildId].ReactionMessage = messageId
	if err := Save(guildId); err != nil {
		log.Errorf("unable to save config file (%v) => %v\n", guildId, err)
	}
}

func GetReactionMessageChannel(guildId string) string {
	return configFiles[guildId].ReactionMessage
}
