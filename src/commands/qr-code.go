package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"regexp"
	"time"
)

func findChannelFromArgs(msg *MessageParser) (*discordgo.Channel, error) {
	var channelId string
	for _, i := range msg.args {
		if matched, err := regexp.Match("<#\\d{18}>", []byte(i)); err != nil {
			skeleton.Bot.SendInternalServerErrorMessage(msg.channel)
			return nil, err
		} else if matched {
			channelId = i
		}
		skeleton.Bot.Logf("arg: %v\n", i)
	}
	if channelId == "" {
		skeleton.Bot.Logf("No channel id detected from %+v\n", msg.args)
		skeleton.Bot.SendMessage(msg.channel, "Aucun channel n'a été détecté")
		return nil, nil
	}
	channelData, err := skeleton.Bot.GetSession().Channel(channelId[2 : len(channelId)-1])
	if err != nil {
		skeleton.Bot.Logf("Error while getting channel '%+v'\n", err)
		if err.Error() == "HTTP 403 Forbidden, {\"message\": \"Missing Access\", \"code\": 50001}" {
			skeleton.Bot.SendMessage(msg.channel, "Je n'ai pas accés à ce channel")
			return nil, err
		}
		skeleton.Bot.SendInternalServerErrorMessage(msg.channel)
		return nil, err
	}
	return channelData, nil
}

func SendDMToAuthor(invitationCode, fullChannelName, recipientID string) error {
	if recipientChannel, err := skeleton.Bot.GetSession().UserChannelCreate(recipientID); err != nil {
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
			Description: "Ce QR code est un lien d'invitation vers le channel " + fullChannelName + ".",
			Timestamp:   time.Now().Format(time.RFC3339),
		}); err != nil {
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
		Guild:                    skeleton.Bot.GetGuildDataFromId(guildId),
		Channel:                  channel,
		Inviter:                  skeleton.Bot.GetUser(),
		CreatedAt:                discordgo.Timestamp(time.Now().String()),
		MaxAge:                   0,
		Revoked:                  false,
		Temporary:                false,
		Unique:                   false,
		TargetUser:               nil,
		TargetUserType:           0,
		ApproximatePresenceCount: 0,
		ApproximateMemberCount:   0,
	}); err != nil {
		return nil, err
	} else {
		return invite, nil
	}
}

func deleteLastMessage(channelId, messageId string) error {
	return skeleton.Bot.GetSession().ChannelMessageDelete(channelId, messageId)
}

func GenerateQrCode(msg *MessageParser) error {
	channel, err := findChannelFromArgs(msg)
	if err != nil {
		return err
	} else if channel == nil {
		return nil
	}
	invitation, err := generateInvitation(channel.ID, msg.guildId)
	if err != nil {
		return err
	}
	if err = SendDMToAuthor(invitation.Code, channel.Mention(), msg.author.ID); err != nil {
		return err
	}
	return deleteLastMessage(msg.channel, msg.message.ID)
}