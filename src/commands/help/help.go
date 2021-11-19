package help

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"log"
)

type Command struct {
	commands.Metadata
	listedCmd []commands.Metadata
}

func (Command) GetName() string {
	return "help"
}

func (Command) GetDescription() string {
	return "Get a list of all commands with their respective description"
}

func (Command) GetOptions() []*discordgo.ApplicationCommandOption {
	return nil
}

func (c *Command) AddCommand(cmd commands.Metadata) {
	c.listedCmd = append(c.listedCmd, cmd)
}

func (c Command) GetHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	commandsList := ""
	for _, cmd := range c.listedCmd {
		commandsList += fmt.Sprintf("**%v**: %v\n", cmd.GetName(), cmd.GetDescription())
	}
	commandsList = commandsList[:len(commandsList)-1]
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Type:        discordgo.EmbedTypeRich,
					Title:       "Voici la liste des commandes disponibles",
					Description: commandsList,
					Color:       skeleton.GenerateRandomMessageColor(),
				},
			},
		},
	})
	if err != nil {
		skeleton.Bot.Logf("help interaction responded with an error: %v", err)
		return
	}
}

type Registerer struct {
	commands.Registerer
	ownCommand []commands.Metadata
}

// NewRegister create a register with a command that must be the help command with
// others commands registered
func NewRegister(c *Command) Registerer {
	return Registerer{
		ownCommand: []commands.Metadata{c},
	}
}

func (h Registerer) GetCommands() []commands.Metadata {
	if len(h.ownCommand) == 0 {
		log.Fatalf("help registerer is not supposed to be declared without 'NewRegister'")
	}
	return h.ownCommand

}
