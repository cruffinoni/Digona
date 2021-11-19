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

func CreateInteractionResponse(description string, color int, flags uint64) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Type:        discordgo.EmbedTypeRich,
					Description: description,
					Color:       color,
				},
			},
			Flags: flags,
		},
	}
}

func RespondErrorInteractionToUserOnly(session *discordgo.Session, interaction *discordgo.Interaction, errorDescription string) {
	if err := session.InteractionRespond(interaction, CreateInteractionResponse(errorDescription, skeleton.RedColor, 1<<6)); err != nil {
		skeleton.Bot.Logf("err msg interaction responded with an error: %v", err)
	}
}

func RespondErrorInteraction(session *discordgo.Session, interaction *discordgo.Interaction, errorDescription string) {
	if err := session.InteractionRespond(interaction, CreateInteractionResponse(errorDescription, skeleton.RedColor, 0)); err != nil {
		skeleton.Bot.Logf("err msg interaction responded with an error: %v", err)
	}
}

func RespondInteraction(session *discordgo.Session, interaction *discordgo.Interaction, description string) {
	if err := session.InteractionRespond(interaction, CreateInteractionResponse(description, skeleton.GenerateRandomMessageColor(), 0)); err != nil {
		skeleton.Bot.Logf("err msg interaction responded with an error: %v", err)
	}
}
