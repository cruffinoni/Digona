package digona

func (bot BotData) DisplayError(channelId, message string) (err error) {
	_, err = bot.Session.ChannelMessageSend(channelId, message)
	return
}
