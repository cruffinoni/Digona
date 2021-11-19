package commands

import "github.com/bwmarrin/discordgo"

type Metadata interface {
	GetName() string
	GetDescription() string
	GetOptions() []*discordgo.ApplicationCommandOption
	GetHandler(*discordgo.Session, *discordgo.InteractionCreate)
}

type Registerer interface {
	GetCommands() []Metadata
}

func GenerateApplicationCommand(c Metadata) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        c.GetName(),
		Description: c.GetDescription(),
		Options:     c.GetOptions(),
	}
}
