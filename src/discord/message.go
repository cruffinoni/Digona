package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"regexp"
)

func DeleteMessage(channelId, messageId string) error {
	return skeleton.Bot.GetSession().ChannelMessageDelete(channelId, messageId)
}

func FindChanelFromArgs(args []string) (string, error) {
	var channelId string
	for _, i := range args {
		if matched, err := regexp.Match("<#\\d{18}>", []byte(i)); err != nil {
			return "", err
		} else if matched {
			channelId = i
		}
	}
	return channelId, nil
}

func GetChannelDataFromRawId(channelId string) (*discordgo.Channel, error) {
	return skeleton.Bot.GetSession().Channel(channelId[2 : len(channelId)-1])
}
