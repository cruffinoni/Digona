package user

import (
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"math"
)

func AddPoint(guildId, userId string) error {
	return skeleton.Bot.GetDatabase().AddPointForMessage(guildId, userId)
}

// CalculatePointForLevel is an geometric suite
// q = 5
// u0 = 5
// un = u0 * q ^ n
func CalculatePointForLevel(level uint) uint32 {
	return uint32(5 * math.Pow(2.0, float64(level)))
}

func IsPassingLevel(guildId, userId string) bool {
	return StoredUsers[guildId][userId].Point >= CalculatePointForLevel(StoredUsers[guildId][userId].Level)
}

func SetNewLevel(guildId, userId string) error {
	for IsPassingLevel(guildId, userId) {
		StoredUsers[guildId][userId].Level++
	}
	return skeleton.Bot.GetDatabase().SetLevel(*StoredUsers[guildId][userId])
}
