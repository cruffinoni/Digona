package dispatch

import (
	"fmt"
	"github.com/Digona/commands"
	"github.com/Digona/digona"
	"github.com/bwmarrin/discordgo"
	"log"
)

func OnMessageCreated(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == digona.Bot.GetID() {
		return
	}
	//fmt.Printf("A message has been created %+#v\n", *message)
	//fmt.Printf("Message: '%v'\n", message.Content)
	//fmt.Printf("Mentions role: '%v'\n", message.MentionRoles)
	//fmt.Printf("Mentions: '%v'\n", message.Mentions)
	if message.ChannelID != "550772992615907328" {
		fmt.Printf("I'm not allowed to read in this channel.\n")
		return
	}
	newParser := commands.New(message)
	//fmt.Printf("The current session: %v\n")
	if !newParser.IsBotMentioned() {
		fmt.Printf("The author might not talk to me.\n")
		return
	}
	if newParser.GetOriginalCommand() == "" {
		_ = digona.Bot.DisplayError(newParser.GetChannelId(), "Aucune commande n'a été détectée! :/\n")
		return
	}
	fmt.Printf("command '%v' detected with args: '%v'\n", newParser.GetOriginalCommand(), newParser.GetArguments())
	if err := newParser.Handler(newParser); err != nil {
		log.Printf("An error occured during executing command '%v' with error '%v'\n", newParser.GetOriginalCommand(), err.Error())
	}
	//_, err := session.ChannelMessageSend(message.ChannelID, "Mentions done =>" + message.Mentions[0].String())
	//if err != nil {
	//	log.Printf("An error occured during sending a message: %v\n", err.Error())
	//}
}
