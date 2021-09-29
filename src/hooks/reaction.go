package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands"
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
	if roleId := commands.GetRoleFromMessageReaction(message.MessageID, reactionId); roleId != "" {
		if err := skeleton.Bot.GetSession().GuildMemberRoleAdd(skeleton.Bot.GetGuildId(), message.UserID, roleId); err != nil {
			skeleton.Bot.Errorf("An error occured while setting the role '%v' to '%v': %v\n", roleId, message.UserID, err)
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
	if roleId := commands.GetRoleFromMessageReaction(message.MessageID, reactionId); roleId != "" {
		if err := skeleton.Bot.GetSession().GuildMemberRoleRemove(skeleton.Bot.GetGuildId(), message.UserID, roleId); err != nil {
			skeleton.Bot.Errorf("An error occured while removing the role '%v' to '%v': %v\n", roleId, message.UserID, err)
		}
	}
}
