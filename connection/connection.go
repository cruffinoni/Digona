package connection

import (
	"fmt"
	"github.com/Digona/digona"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

func getBotToken() (token string) {
	token = os.Getenv("DIGONA_BOT_TOKEN")
	if token == "" {
		log.Fatal("No token set in the environment variable DIGONA_BOT_TOKEN.\n")
	}
	return
}

func GetFormattedToken() string {
	return fmt.Sprintf("Bot %v", getBotToken())
}

func InitBot(bot *digona.BotData) error {
	fmt.Printf("Digona (version: %v), initialization...\n", digona.BotVersion)
	session, err := discordgo.New(GetFormattedToken())
	if err != nil {
		log.Fatalf("An error occured at the bot creation: %v\n", err.Error())
	}
	bot.SetSession(session)
	err = digona.RetrieveBotData()
	if err != nil {
		log.Fatalf("Cannot retrieve own bot's info: %v\n", err.Error())
	}
	bot.UpTime = time.Now().UTC()
	return err
}
