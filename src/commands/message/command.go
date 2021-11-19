package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
)

type command struct {
	commands.Metadata
}

func (command) GetName() string {
	return "delete-message"
}

func (command) GetDescription() string {
	return "Delete X last messages"
}

func (command) GetOptions() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionInteger,
			Name:        "count",
			Description: "The total of messages to delete",
			Required:    true,
		},
	}
}

func (command) GetHandler(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	interactData := interaction.ApplicationCommandData()
	msgToDelete := int(interactData.Options[0].IntValue())
	if msgToDelete < minAmountDeleteMsg || msgToDelete > (maxAmountDeleteMsg-1) {
		skeleton.Bot.SendMessageWithNoTitle(interaction.ChannelID, fmt.Sprintf("Je ne peux supprimer qu'entre %v et %v messages à la fois.", minAmountDeleteMsg, maxAmountDeleteMsg-1))
		return
	}

	fakeMsg, err := s.ChannelMessageSend(interaction.ChannelID, "fake msg")
	if err != nil {
		discord.RespondErrorInteraction(s, interaction.Interaction, "Je n'ai pas les droits de poster un message pour supprimer ceux précédents")
		return
	}
	var allMessages []*discordgo.Message
	allMessages, err = skeleton.Bot.GetSession().ChannelMessages(interaction.ChannelID, msgToDelete+1, "", "", fakeMsg.ID)
	if err != nil {
		discord.RespondErrorInteraction(s, interaction.Interaction, "Je n'arrive pas à supprimer les messages... Essayez dans quelques minutes.")
		return
	}
	var msgsDeleted uint
	msgsDeleted, err = deleteLastMessages(interaction.ChannelID, allMessages)
	if msgsDeleted > 0 {
		msgsDeleted--
	}
	if err != nil {
		discord.RespondErrorInteractionToUserOnly(s, interaction.Interaction, fmt.Sprintf("Seulement %v messages (sur %v) ont été supprimés", msgsDeleted, msgToDelete))
		skeleton.Bot.Errorf("Err while deleted last messages: %v\n", err)
	}
	discord.RespondInteraction(s, interaction.Interaction, fmt.Sprintf("%v message ont été supprimés", msgsDeleted))
}

type Registerer struct {
	commands.Registerer
}

func (Registerer) GetCommands() []commands.Metadata {
	return []commands.Metadata{
		command{},
	}
}
