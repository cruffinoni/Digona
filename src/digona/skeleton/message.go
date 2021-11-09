package skeleton

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

const (
	TimeoutSec = 10
	redColor   = 15548997
)

func (bot BotData) sendMessageToChannel(channelId, messageContent, title string, delayed, forceDescription bool, color int) {
	var (
		footerText string
		content    string
	)
	if delayed {
		footerText = fmt.Sprintf("Ce message sera supprimé dans %v sec.", TimeoutSec)
	}
	if color == 0 {
		color = GenerateRandomMessageColor()
	}
	if len(messageContent) >= 0xFF || forceDescription {
		content = messageContent
	} else {
		title = messageContent
	}
	if message, err := bot.session.ChannelMessageSendEmbed(channelId, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       title,
		Description: content,
		Color:       color,
		Footer: &discordgo.MessageEmbedFooter{
			Text: footerText,
		},
	}); err != nil {
		bot.Errorf("Can't display a message into the channel '%v'. Error: '%v'\n", channelId, err)
	} else if delayed {
		go func() {
			time.Sleep(time.Second * TimeoutSec)
			if err = bot.session.ChannelMessageDelete(channelId, message.ID); err != nil {
				bot.Errorf("Can't remove a message previously send. Error: '%v'\n", channelId, err)
			}
		}()
	}
}

func (bot BotData) SendMessageWithTitle(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, "", false, false, 0)
}

func (bot BotData) SendMessageWithNoTitle(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, "", false, true, 0)
}

func (bot BotData) SendErrorMessage(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, "", false, false, redColor)
}

func (bot BotData) SendDelayedErrorMessage(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, "", true, false, redColor)
}

func (bot BotData) SendInternalServerErrorMessage(channelId string) {
	bot.sendMessageToChannel(channelId, "", "Une erreur s'est produite, réessayez plus tard", false, false, redColor)
}

func (bot BotData) SendDelayedInternalServerErrorMessage(channelId string) {
	bot.sendMessageToChannel(channelId, "", "Une erreur s'est produite, réessayez plus tard", true, false, redColor)
}

func (bot BotData) SendDelayedMessage(channelId, message string) {
	bot.sendMessageToChannel(channelId, message, "", true, false, 0)
}
