package handler

import (
	"fmt"
	"github.com/Digona/commands/dispatch"
	"github.com/bwmarrin/discordgo"
)

var discordHandlers = []interface{}{
	onBotReady,
	dispatch.OnMessageCreated,
}

func RegisterHandler(session *discordgo.Session) {
	fmt.Print("Creating all discord handler..\n")
	for _, function := range discordHandlers {
		session.AddHandler(function)
	}
	fmt.Print("All handler has been successfully created.\n")
}

func onBotReady(session *discordgo.Session, ready *discordgo.Ready) {
	fmt.Printf("Digona has been succefully initialized and ready!\n")
}
