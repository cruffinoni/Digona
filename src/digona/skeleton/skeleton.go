package skeleton

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/logger"
	"time"
)

type BotData struct {
	session   *discordgo.Session
	data      *discordgo.User
	guild     map[string]*discordgo.Guild
	startTime time.Time
	logger.Logger
}

var Bot BotData

const BotVersion = "0.0.5"

func (bot *BotData) RetrieveInfo() (err error) {
	bot.data, err = bot.session.User("@me")
	if err != nil {
		return err
	}
	bot.guild = make(map[string]*discordgo.Guild)
	return err
}

func (bot *BotData) RegisterGuild(guild *discordgo.Guild) {
	bot.guild[guild.ID] = guild
}

func (bot BotData) GetSession() *discordgo.Session {
	return bot.session
}

func (bot *BotData) SetSession(session *discordgo.Session) {
	if bot.session != nil {
		return
	}
	bot.session = session
}

func (bot BotData) GetID() string {
	return bot.data.ID
}

func (bot BotData) GetUser() *discordgo.User {
	return bot.data
}

func (bot BotData) GetMention() string {
	return fmt.Sprintf("<@!%v>", bot.data.ID)
}

func (bot BotData) GetGuildDataFromId(guildId string) *discordgo.Guild {
	if data, ok := bot.guild[guildId]; !ok {
		bot.Errorf("Can't get data from guild because it's an invalid id: %v\n", guildId)
		return nil
	} else {
		return data
	}
}

func (bot *BotData) StartTime() {
	bot.startTime = time.Now().UTC()
}
