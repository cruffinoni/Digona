package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/handler"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

func OnMessageCreated(_ *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == skeleton.Bot.GetID() {
		return
	}
	skeleton.Bot.Logf("> [%v] %v\n", message.Author.Username, message.ContentWithMentionsReplaced())
	newParser := parser.New(message, skeleton.Bot.Logger)
	if !newParser.IsBotMentioned() {
		skeleton.Bot.Log("The author might not talk to me.\n")
		return
	}
	cmd := handler.GetCommandFromArgs(newParser.GetArguments())
	if cmd == nil {
		skeleton.Bot.Log("No command detected\n")
		return
	}
	newParser.RemoveArgument(cmd.Name)
	if err := cmd.Command(newParser); err != nil {
		skeleton.Bot.SendInternalServerErrorMessage(message.ChannelID)
		skeleton.Bot.Errorf("An error occurred during executing command '%v' with error '%v'\n", newParser.GetOriginalCommand(), err.Error())
	}
}
