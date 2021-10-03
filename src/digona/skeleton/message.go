package skeleton

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

const timeOutSec = 10

func (bot BotData) sendMessageToChannel(channelId, messageContent string, delayed bool) {
	var footerText string
	if delayed {
		footerText = fmt.Sprintf("Ce message sera supprimé dans %v sec.", timeOutSec)
	}
	if message, err := bot.session.ChannelMessageSendEmbed(channelId, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Description: messageContent,
		Color:       GenerateRandomMessageColor(),
		Footer: &discordgo.MessageEmbedFooter{
			Text: footerText,
		},
	}); err != nil {
		bot.Errorf("Can't display a message into the channel '%v'. Error: '%v'\n", channelId, err)
	} else if delayed {
		go func() {
			time.Sleep(time.Second * timeOutSec)
			if err = bot.session.ChannelMessageDelete(channelId, message.ID); err != nil {
				bot.Errorf("Can't remove a message previously send. Error: '%v'\n", channelId, err)
			}
		}()
	}
}

func (bot BotData) SendMessage(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, false)
}

func (bot BotData) SendInternalServerErrorMessage(channelId string) {
	bot.sendMessageToChannel(channelId, "Une erreur s'est produite, réessayez plus tard", false)
}

func (bot BotData) SendInternalServerErrorMessageTimeout(channelId string) {
	bot.sendMessageToChannel(channelId, "Une erreur s'est produite, réessayez plus tard", true)
}

func (bot BotData) SendDelayedMessage(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, true)
}
