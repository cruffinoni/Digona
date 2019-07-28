package digona

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type BotData struct {
	Session *discordgo.Session
	data *discordgo.User
	UpTime time.Time
	Game discordgo.Game
}

var Bot BotData

const DigonaVersion = "0.0.1"

func RetrieveBotData() (err error) {
	Bot.data, err = Bot.Session.User("@me")
	return err
}

func (bot BotData) GetID() string {
	return bot.data.ID
}