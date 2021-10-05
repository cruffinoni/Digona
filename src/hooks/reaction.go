package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/reaction"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

func OnUserReact(_ *discordgo.Session, message *discordgo.MessageReactionAdd) {
	if message.UserID == skeleton.Bot.GetID() {
		return
	}
	var reactionId string
	if message.Emoji.ID != "" {
		reactionId = message.Emoji.APIName()
	} else {
		reactionId = message.Emoji.Name
	}
	if roleId := reaction.GetRoleFromMessageReaction(message.MessageID, reactionId); roleId != "" {
		if err := skeleton.Bot.GetSession().GuildMemberRoleAdd(message.GuildID, message.UserID, roleId); err != nil {
			skeleton.Bot.Errorf("An error occurred while setting the role '%v' to '%v': %v\n", roleId, message.UserID, err)
		}
	}
}

func OnUserRemoveReact(_ *discordgo.Session, message *discordgo.MessageReactionRemove) {
	if message.UserID == skeleton.Bot.GetID() {
		return
	}
	var reactionId string
	if message.Emoji.ID != "" {
		reactionId = message.Emoji.APIName()
	} else {
		reactionId = message.Emoji.Name
	}
	if roleId := reaction.GetRoleFromMessageReaction(message.MessageID, reactionId); roleId != "" {
		if err := skeleton.Bot.GetSession().GuildMemberRoleRemove(message.GuildID, message.UserID, roleId); err != nil {
			skeleton.Bot.Errorf("An error occurred while removing the role '%v' to '%v': %v\n", roleId, message.UserID, err)
		}
	}
}
