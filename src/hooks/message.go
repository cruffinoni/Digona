package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/handler"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/logger"
)

func OnMessageCreated(_ *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == skeleton.Bot.GetID() {
		return
	}
	log := logger.Logger{}
	log.Logf("> [%v] %v\n", message.Author.Username, message.ContentWithMentionsReplaced())
	newParser := parser.New(message, log)
	if !newParser.IsBotMentioned() {
		log.Log("The author might not talk to me.\n")
		return
	}
	cmd := handler.GetCommandFromArgs(newParser.GetArguments())
	if cmd == nil {
		log.Log("No command detected\n")
		return
	}
	newParser.RemoveArgument(cmd.Name)
	if err := cmd.Command(newParser); err != nil {
		skeleton.Bot.SendInternalServerErrorMessage(message.ChannelID)
		log.Errorf("An error occurred during executing command '%v' with error '%v'\n", newParser.GetOriginalCommand(), err.Error())
	}
}
