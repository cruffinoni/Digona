package user

import (
	"github.com/cruffinoni/Digona/src/database/models"
)

type User struct {
	DatabaseModel models.TableUser
}

type userMap map[string]*models.TableUser

var (
	StoredUsers = make(map[string]userMap)
)

func StoreUsersFromModels(guildId string, model []models.TableUser) {
	if StoredUsers[guildId] == nil {
		StoredUsers[guildId] = make(map[string]*models.TableUser)
	}
	for i := range model {
		StoredUsers[guildId][model[i].DiscordId] = &model[i]
	}
}
