package models

import "gorm.io/gorm"

type TableUser struct {
	gorm.Model
	DiscordId string `gorm:"uniqueIndex:idx_unique_discord_guild"`
	GuildId   string `gorm:"uniqueIndex:idx_unique_discord_guild"`
	Point     int
}
