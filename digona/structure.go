package digona

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"time"
)

type BotData struct {
	session *discordgo.Session
	data    *discordgo.User
	UpTime  time.Time
	Game    discordgo.Game
}

var Bot BotData

const BotVersion = "0.0.1"

func RetrieveBotData() (err error) {
	Bot.data, err = Bot.session.User("@me")
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
