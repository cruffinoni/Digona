package user

import "github.com/cruffinoni/Digona/src/digona/skeleton"

func AddPoint(guildId, userId string) error {
	return skeleton.Bot.GetDatabase().AddPointForMessage(guildId, userId)
}
