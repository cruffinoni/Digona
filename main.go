package main

import (
	"github.com/cruffinoni/Digona/src/digona"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	digona.Init(&skeleton.Bot)
	botHandler := make(chan os.Signal, 1)
	signal.Notify(botHandler, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-botHandler
	if err := skeleton.Bot.GetSession().Close(); err != nil {
		skeleton.Bot.Fatalf("Error occurred during the bot disconnection: %v\n", err.Error())
	}
}
