package skeleton

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/logger"
	"os"
	"time"
)

type BotData struct {
	session   *discordgo.Session
	data      *discordgo.User
	guild     *discordgo.Guild
	startTime time.Time
	logger.Logger
}

var Bot BotData

const BotVersion = "0.0.3"

func (bot *BotData) RetrieveInfo() (err error) {
	bot.data, err = bot.session.User("@me")
	if err != nil {
		return err
	}
	guildId := os.Getenv("GUILD_ID")
	if guildId != "" {
		bot.guild, err = bot.session.Guild(guildId)
	}
	return err
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

func (bot BotData) GetMention() string {
	return fmt.Sprintf("<@!%v>", bot.data.ID)
}

func (bot BotData) GetGuildId() string {
	return bot.guild.ID
}

func (bot *BotData) StartTime() {
	bot.startTime = time.Now().UTC()
}
