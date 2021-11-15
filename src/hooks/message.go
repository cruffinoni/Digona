package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/handler"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"strings"
)

func logCmdError(channelId, cmdName string, err error) {
	skeleton.Bot.SendInternalServerErrorMessage(channelId)
	skeleton.Bot.Errorf("An error occurred during executing command '%v' with error '%v'\n", cmdName, err.Error())
}

func isHiMessagePresent(words []string) bool {
	validHiMessage := []string{
		"hey", "hi", "salut", "hello", "wsh", "wesh", "bonjour", "bonsoir",
	}
	for _, w := range words {
		for _, vHi := range validHiMessage {
			if strings.ToLower(w) == vHi {
				return true
			}
		}
	}
	return false
}

func OnMessageCreated(_ *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == skeleton.Bot.GetID() {
		return
	}
	var err error
	skeleton.Bot.Logf("> [%v] [%v]: %v\n", skeleton.Bot.GetGuildDataFromId(message.GuildID).Name, message.Author.Username, message.ContentWithMentionsReplaced())
	newParser := parser.New(message, skeleton.Bot.Logger)
	if !newParser.IsBotMentioned() {
		if err = skeleton.Bot.GetDatabase().AddPointForMessage(message.GuildID, message.Author.ID); err != nil {
			skeleton.Bot.Errorf("can't add point for new message: '%v'\n", err.Error())
			return
		}
		if isHiMessagePresent(newParser.GetArguments()) {
			if err = skeleton.Bot.GetSession().MessageReactionAdd(message.ChannelID, message.ID, `👋`); err != nil {
				skeleton.Bot.Errorf("Can't add a reaction to the 'hi' message: %v", err)
			}
		}
		skeleton.Bot.Log("The author might not talk to me.\n")
		return
	}
	if newParser.GetArguments()[0] == "help" {
		if err = handler.HelpCommand(newParser); err != nil {
			logCmdError(message.ChannelID, "help", err)
			return
		}
	}
	cmd := handler.GetCommandFromArgs(newParser.GetArguments())
	if cmd == nil {
		skeleton.Bot.Log("No command detected\n")
		return
	}
	newParser.RemoveArgument(cmd.Name)
	if err = cmd.Command(newParser); err != nil {
		logCmdError(message.ChannelID, newParser.GetOriginalCommand(), err)
		return
	}
}
