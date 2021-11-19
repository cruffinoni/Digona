package config

import (
	"encoding/json"
	"github.com/cruffinoni/Digona/src/database/models"
	"github.com/cruffinoni/Digona/src/digona/version"
	"github.com/cruffinoni/Digona/src/logger"
)

var (
	configFiles = make(map[string]*models.ConfigFileHolder)
	log         logger.Logger
)

func GenerateConfigFileHolder() *models.ConfigFileHolder {
	return &models.ConfigFileHolder{
		Version: version.BotVersion,
	}
}

func StoreConfig(guildId string, config *models.ConfigFileHolder) {
	configFiles[guildId] = config
}

func StoreConfigFiles(configs []models.TableConfig) error {
	var (
		localJson models.ConfigFileHolder
		err       error
	)
	for _, c := range configs {
		if err = json.Unmarshal(c.Config, &localJson); err != nil {
			return err
		}
		configFiles[c.GuildId] = &localJson
	}
	return nil
}

func DoesExists(guildId string) bool {
	_, exists := configFiles[guildId]
	return exists
}
