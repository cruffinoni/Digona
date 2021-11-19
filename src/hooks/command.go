package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/handler"
	"log"
)

func OnUserUseSlashCommands(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if cmdHandler := handler.GetCommandHandler(interaction.ApplicationCommandData().Name); cmdHandler != nil {
		cmdHandler(session, interaction)
	} else {
		log.Printf("Interaction handler not implemented")
	}
}
