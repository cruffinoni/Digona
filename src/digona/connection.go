package digona

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/config"
	"github.com/cruffinoni/Digona/src/database/models"
	"github.com/cruffinoni/Digona/src/digona/handler"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/digona/version"
	"log"
	"os"
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

func Init(bot *skeleton.BotData) {
	bot.Logf("Digona (version: %v), initialization...\n", version.BotVersion)
	discordgo.MakeIntent(discordgo.IntentsAll)
	bot.InitDatabase()
	session, err := discordgo.New(GetFormattedToken())
	if err != nil {
		bot.Fatalf("An error occurred at the bot creation: %v\n", err.Error())
	}
	bot.Log("Loading config files...\n")
	var configs []models.TableConfig
	if configs, err = bot.GetDatabase().LoadAllConfigFiles(); err != nil {
		bot.Fatalf("error while loading the config files: %v", err)
	}
	if err = config.StoreConfigFiles(configs); err != nil {
		bot.Fatalf("can't store the config files: %v", err)
	}
	bot.Log("Setting Discord session...\n")
	bot.SetSession(session)
	bot.Log("Retrieving bot's infos...\n")
	if err = bot.RetrieveInfo(); err != nil {
		bot.Fatalf("Cannot retrieve own bot infos: %v\n", err.Error())
	}
	handler.RegisterHandler(session)
	bot.Log("Registering hooks to discord...\n")
	if err = session.Open(); err != nil {
		bot.Fatalf("Can't open the session: %v\n", err.Error())
	}
}
