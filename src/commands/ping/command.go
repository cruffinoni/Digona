package ping

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

type command struct {
	commands.Metadata
}

func (command) GetName() string {
	return "ping"
}

func (command) GetDescription() string {
	return "A simple command that returns 'pong'"
}

func (command) GetOptions() []*discordgo.ApplicationCommandOption {
	return nil
}

func (command) GetHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Type:        discordgo.EmbedTypeRich,
					Description: "Pong!",
					Color:       skeleton.GenerateRandomMessageColor(),
				},
			},
		},
	})
	if err != nil {
		skeleton.Bot.Logf("ping interaction responded with an error: %v", err)
	}
}

type Registerer struct {
	commands.Registerer
}

func (Registerer) GetCommands() []commands.Metadata {
	return []commands.Metadata{
		command{},
	}
}
