package config

import (
	"bytes"
	"encoding/json"
	"github.com/cruffinoni/Digona/src/digona/version"
	"github.com/cruffinoni/Digona/src/logger"
	"io/ioutil"
	"os"
	"strings"
)

var (
	configFiles = make(map[string]*ConfigurationFile)
	log         logger.Logger
)

type Configuration struct {
	WelcomeChannel  string `json:"welcome_channel"`
	ReactionMessage string `json:"reaction_message"`
}

type ConfigurationFile struct {
	Version       string `json:"version"`
	Configuration `json:"configuration"`
}

func FileExists(guildId string) bool {
	_, exists := configFiles[guildId]
	return exists
}

func Create(guildId string) error {
	if FileExists(guildId) {
		log.Logf("Config file for %v already exists\n", guildId)
		return nil
	}
	configFiles[guildId] = &ConfigurationFile{
		Version:       version.BotVersion,
		Configuration: Configuration{},
	}
	return Save(guildId)
}

func formatJSONFileName(guildId string) string {
	return "./config/" + guildId + ".json"
}

func Save(guildId string) error {
	if file, err := os.Create(formatJSONFileName(guildId)); err != nil {
		log.Logf("Can't save file for guild id %v: %v\n", guildId, err)
		return err
	} else {
		reqBodyBytes := new(bytes.Buffer)
		if err = json.NewEncoder(reqBodyBytes).Encode(configFiles[guildId]); err != nil {
			log.Logf("can't encode json to bytes (%v): %v\n", guildId, err)
			return err
		}
		if _, err = file.Write(reqBodyBytes.Bytes()); err != nil {
			log.Logf("can't write to a new file (%v): %v\n", guildId, err)
			return err
		}
		return nil
	}
}

func Load() {
	files, err := ioutil.ReadDir("./config/")
	if err != nil {
		log.FatalMsg(err)
	}
	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".json") {
			continue
		}
		fileContent, err := os.Open("./config/" + file.Name())
		if err != nil {
			log.FatalMsg(err)
		}
		byteValue, err := ioutil.ReadAll(fileContent)
		if err != nil {
			log.Fatalf("Can't read content file: %v\n", err)
		}
		guildId := file.Name()[:len(file.Name())-4]
		config := ConfigurationFile{}
		if err := json.Unmarshal(byteValue, &config); err != nil {
			log.Fatalf("Cant unmarshal data from json (%v): %v\n", guildId, err)
		}
		configFiles[guildId] = &config
	}
	log.Logf("%v config files loaded\n", len(configFiles))
}
