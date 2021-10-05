package privateMessage

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"time"
)

func findChannelFromArgs(msg *parser.MessageParser) (*discordgo.Channel, error) {
	channelId, err := discord.FindChanelFromArgs(msg.GetArguments())
	if err != nil {
		skeleton.Bot.Logf("Error from discord %+v\n", msg.GetArguments())
		skeleton.Bot.SendInternalServerErrorMessage(msg.GetChannelId())
		return nil, err
	} else if channelId == "" {
		skeleton.Bot.Logf("No channel id detected from %+v\n", msg.GetArguments())
		skeleton.Bot.SendMessage(msg.GetChannelId(), "Aucun channel n'a été détecté")
		return nil, nil
	}
	channelData, err := discord.GetChannelDataFromRawId(channelId)
	if err != nil {
		skeleton.Bot.Logf("Error while getting channel '%+v'\n", err)
		if discord.IsMissingAccessError(err) {
			skeleton.Bot.SendMessage(msg.GetChannelId(), "Je n'ai pas accés à ce channel")
			return nil, err
		}
		skeleton.Bot.SendInternalServerErrorMessage(msg.GetChannelId())
		return nil, err
	}
	return channelData, nil
}

func SendDMToAuthor(invitationCode, channelMention, recipientID, originalChannelId string) error {
	if recipientChannel, err := skeleton.Bot.GetSession().UserChannelCreate(recipientID); err != nil {
		skeleton.Bot.SendInternalServerErrorMessage(originalChannelId)
		skeleton.Bot.Logf("Unable to create a private message w/ %v => %v\n", recipientID, err.Error())
		return err
	} else {
		URL := "https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=https://discord.gg/" + invitationCode
		if _, err = skeleton.Bot.GetSession().ChannelMessageSendEmbed(recipientChannel.ID, &discordgo.MessageEmbed{
			Type:  discordgo.EmbedTypeImage,
			Color: skeleton.GenerateRandomMessageColor(),
			Image: &discordgo.MessageEmbedImage{
				URL:    URL,
				Width:  150,
				Height: 150,
			},
			Title:       "QR Code",
			Description: "Ce QR code est un lien d'invitation vers le channel " + channelMention + ".",
			Timestamp:   time.Now().Format(time.RFC3339),
		}); err != nil {
			skeleton.Bot.SendInternalServerErrorMessage(originalChannelId)
			skeleton.Bot.Logf("Can't send a private url to %v => %v\n", recipientID, err.Error())
			return err
		}
	}
	return nil
}

func generateInvitation(channelId, guildId string) (*discordgo.Invite, error) {
	channel, err := skeleton.Bot.GetSession().Channel(channelId)
	if err != nil {
		return nil, err
	}
	if invite, err := skeleton.Bot.GetSession().ChannelInviteCreate(channelId, discordgo.Invite{
		Guild:     skeleton.Bot.GetGuildDataFromId(guildId),
		Channel:   channel,
		Inviter:   skeleton.Bot.GetUser(),
		CreatedAt: discordgo.Timestamp(time.Now().String()),
		MaxAge:    0,
		Revoked:   false,
		Temporary: false,
		Unique:    true,
	}); err != nil {
		return nil, err
	} else {
		return invite, nil
	}
}

func GenerateQrCode(msg *parser.MessageParser) error {
	channel, err := findChannelFromArgs(msg)
	if err != nil {
		return err
	} else if channel == nil {
		return nil
	}
	invitation, err := generateInvitation(channel.ID, msg.GetGuildId())
	if err != nil {
		return err
	}
	if err = SendDMToAuthor(invitation.Code, channel.Mention(), msg.GetDiscordAuthor().ID, msg.GetChannelId()); err != nil {
		return err
	}
	return discord.DeleteMessage(msg.GetChannelId(), msg.GetDiscordMessage().ID)
}
