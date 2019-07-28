package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type botStructure struct {
	session *discordgo.Session
	data *discordgo.User
	upTime time.Time
	game discordgo.Game
}

var DigonaBot botStructure

const digonaVersion = "0.0.1"

func main() {
	var err error
	if initBot() != nil {
		return
	}
	RegisterHandler(DigonaBot.session)
	err = DigonaBot.session.Open()
	if err != nil {
		log.Printf("Error occured during the bot connection: %v\n", err.Error())
		os.Exit(1)
	}
	botHandler := make(chan os.Signal, 1)
	signal.Notify(botHandler, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-botHandler
	err = DigonaBot.session.Close()
	if err != nil {
		log.Printf("Error occured during the bot deconnection: %v\n", err.Error())
	}
}