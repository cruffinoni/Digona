package digona

func (bot BotData) DisplayError(channelId, message string) (err error) {
	_, err = bot.GetSession().ChannelMessageSend(channelId, message)
	return
}
