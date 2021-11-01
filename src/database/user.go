package database

import (
	"github.com/cruffinoni/Digona/src/database/models"
	"gorm.io/gorm/clause"
)

func (db Database) AddUser(guildId, userId string) (models.TableUser, error) {
	user := models.TableUser{
		DiscordId: userId,
		GuildId:   guildId,
	}
	err := db.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&user).Error
	return user, err
}

func (db Database) LoadUsersForGuild(guildId string) ([]models.TableUser, error) {
	var users []models.TableUser
	err := db.db.Model(&models.TableUser{}).Find(&users, "guild_id = ?", guildId).Error
	return users, err
}

func (db Database) DeleteUsersFromGuild(guildId string) error {
	return db.db.Delete(&models.TableUser{}, "guildId = ?", guildId).Error
}

func (db Database) AddPointForMessage(guildId, userId string) error {
	fullUser := models.TableUser{}
	if err := db.db.Model(&models.TableUser{}).Where("discord_id = ? and guild_id = ?", userId, guildId).Find(&fullUser).Error; err != nil {
		return err
	}
	return db.db.Model(&fullUser).Update("Point", fullUser.Point+5).Error
}

func (db Database) GetBestRankedPlayers(guildId string, max int) (users []models.TableUser, err error) {
	err = db.db.Model(&models.TableUser{}).Limit(max).Order("point DESC").Find(&users, "guild_id = ?", guildId).Error
	return
}
