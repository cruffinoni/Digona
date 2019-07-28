package commands

import (
	"fmt"
	"github.com/Digona/src/digona"
	"github.com/bwmarrin/discordgo"
	"log"
)

type commandHandler func(*MessageParser) error

var userCommands = map[string] commandHandler {
	"delete": deleteLastMessages,
}

func OnMessageCreated(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == digona.Bot.GetID() {
		return
	}
	//fmt.Printf("A message has been created %+#v\n", *message)
	//fmt.Printf("Message: '%v'\n", message.Content)
	//fmt.Printf("Mentions role: '%v'\n", message.MentionRoles)
	//fmt.Printf("Mentions: '%v'\n", message.Mentions)
	parser := New(message)
	if !parser.isMentioned {
		fmt.Printf("The author might not talk to me.\n")
		return
	}
	if parser.Command == "" {
		_ = digona.Bot.DisplayError(parser.channel, "Je n'ai pas compris! :/\n")
		return
	}
	fmt.Printf("Command '%v' detected with args: '%v'\n", parser.Command, parser.Args)
	if err := parser.handler(parser); err != nil {
		log.Printf("An error occured during executing command '%v' with error '%v'\n", parser.Command, err.Error())
	}
	//_, err := session.ChannelMessageSend(message.ChannelID, "Mentions done =>" + message.Mentions[0].String())
	//if err != nil {
	//	log.Printf("An error occured during sending a message: %v\n", err.Error())
	//}
}