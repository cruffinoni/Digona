package discord

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func GetReactionFromRawReactionId(reactionId string) string {
	if !strings.Contains(reactionId, ":") {
		return reactionId
	}
	return reactionId[1 : len(reactionId)-1]
}

func GetReactionIdFromRawReactionId(reactionId string) string {
	if !strings.Contains(reactionId, ":") {
		return reactionId
	}
	return reactionId[2 : len(reactionId)-1]
}

func FindEmojiFromRawEmojiId(customEmojis []*discordgo.Emoji, reference string) *discordgo.Emoji {
	reference = GetReactionFromRawReactionId(reference)
	for _, r := range customEmojis {
		fullRef := ":" + r.Name + ":" + r.ID
		if fullRef == reference {
			return r
		}
	}
	return nil
}
