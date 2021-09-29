package skeleton

func (bot BotData) SendMessage(channelId, message string) {
	if _, err := bot.session.ChannelMessageSend(channelId, message); err != nil {
		bot.Errorf("Can't display a message into the channel '%v'. Error: '%v'\n", channelId, err)
	}
}
