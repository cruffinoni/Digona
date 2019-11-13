package main

import (
	"github.com/Digona/connection"
	"github.com/Digona/digona"
	"github.com/Digona/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var err error
	if connection.InitBot(&digona.Bot) != nil {
		return
	}
	handler.RegisterHandler(digona.Bot.GetSession())
	err = digona.Bot.GetSession().Open()
	if err != nil {
		log.Fatalf("Error occured during the bot connection: %v\n", err.Error())
	}
	botHandler := make(chan os.Signal, 1)
	signal.Notify(botHandler, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-botHandler
	err = digona.Bot.GetSession().Close()
	if err != nil {
		log.Fatalf("Error occured during the bot deconnection: %v\n", err.Error())
	}
}
