package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var discordHandlers = []interface{} {
	onBotReady,
	onMessageCreated,
}

func RegisterHandler(session *discordgo.Session) {
	fmt.Print("Creating all discord handler..\n")
	for _, funct := range discordHandlers {
		session.AddHandler(funct)
		fmt.Printf("Handler %v created\n", funct)
	}
	fmt.Print("All handler has been successfully created.\n")
}

func onBotReady(session *discordgo.Session, ready *discordgo.Ready) {
	fmt.Printf("Digona has been succefully connected.\n")
}

func onMessageCreated(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == DigonaBot.data.ID {
		return
	}
	fmt.Printf("A message has been created %+#v\n", *message)
	fmt.Printf("Message: '%v'\n", message.Content)
	fmt.Printf("Mentions role: '%v'\n", message.MentionRoles)
	fmt.Printf("Mentions: '%v'\n", message.Mentions)
	_, err := session.ChannelMessageSend(message.ChannelID, "Mentions done =>" + message.Mentions[0].String())
	if err != nil {
		log.Printf("An error occured during sending a message: %v\n", err.Error())
	}
}