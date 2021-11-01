package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

func GetAllUsersFromGuild(guildId string) ([]*discordgo.Member, error) {
	if users, err := skeleton.Bot.GetSession().GuildMembers(guildId, "", 1000); err != nil {
		return nil, err
	} else {
		return users, err
	}
}
