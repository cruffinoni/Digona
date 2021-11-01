package user

import "github.com/cruffinoni/Digona/src/database/models"

type User struct {
	DatabaseModel models.TableUser
	DiscordId     string
	Point         int
}

var (
	StoredUsers = make(map[string]map[string]*User)
)

func StoreUsersFromModels(guildId string, model []models.TableUser) {
	if StoredUsers[guildId] == nil {
		StoredUsers[guildId] = make(map[string]*User)
	}
	for _, u := range model {
		StoredUsers[guildId][u.DiscordId] = &User{
			DatabaseModel: u,
			DiscordId:     u.DiscordId,
			Point:         u.Point,
		}
	}
}
