package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands"
	"github.com/cruffinoni/Digona/src/commands/help"
	"github.com/cruffinoni/Digona/src/commands/message"
	"github.com/cruffinoni/Digona/src/commands/opgg"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/commands/ping"
	"github.com/cruffinoni/Digona/src/commands/privateMessage"
	"github.com/cruffinoni/Digona/src/commands/reaction"
	"github.com/cruffinoni/Digona/src/commands/role"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"log"
)

type CommandHandler func(*parser.MessageParser) error
type CmdHandler func(*discordgo.Session, *discordgo.InteractionCreate)

type CommandPair struct {
	Name    string
	Command CommandHandler
}

var (
	commandsListing = map[string]CommandHandler{
		//"delete":       message.RedirectDelete,
		"react":        reaction.Role,
		"react-add":    reaction.ChangeReaction,
		"qr-code":      privateMessage.GenerateQrCode,
		"default-role": role.SetDefaultRole,
		//"ranking":      common_pattern.GetRanking,
		"opgg": opgg.GetOPGGLink,
	}

	listedRegisterer = []commands.Registerer{
		ping.Registerer{},
		message.Registerer{},
	}
)

func RegisterCommands(guildId string) {
	// A pointer is required to have the same instance across the code
	helpCmd := &help.Command{}
	// Register help command manually because this command has specific functions
	// to be fully functional like AddCommand
	listedRegisterer = append(listedRegisterer, help.NewRegister(helpCmd))

	for _, i := range listedRegisterer {
		for _, c := range i.GetCommands() {
			_, err := skeleton.Bot.GetSession().ApplicationCommandCreate(skeleton.Bot.GetSession().State.User.ID, guildId, commands.GenerateApplicationCommand(c))
			if err != nil {
				log.Printf("Err while registering a command: %v", err)
			}
			helpCmd.AddCommand(c)
		}
	}
}

func GetCommandHandler(commandName string) CmdHandler {
	for _, i := range listedRegisterer {
		for _, c := range i.GetCommands() {
			if commandName == c.GetName() {
				return c.GetHandler
			}
		}
	}
	return nil
}
