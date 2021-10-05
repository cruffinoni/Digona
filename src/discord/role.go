package discord

import "github.com/bwmarrin/discordgo"

func FindRoleFromRawRoleId(roles []*discordgo.Role, reference string) *discordgo.Role {
	reference = reference[3 : len(reference)-1]
	for _, r := range roles {
		if r.ID == reference {
			return r
		}
	}
	return nil
}

func FindEmojiFromRawEmojiId(customEmojis []*discordgo.Emoji, reference string) *discordgo.Emoji {
	reference = reference[1 : len(reference)-1]
	for _, r := range customEmojis {
		fullRef := ":" + r.Name + ":" + r.ID
		if fullRef == reference {
			return r
		}
	}
	return nil
}
