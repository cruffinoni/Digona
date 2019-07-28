package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

func getBotToken() (token string) {
	token = os.Getenv("DIGONA_BOT_TOKEN")
	if token == "" {
		log.Fatal("No token set in the environment variable DIGONA_BOT_TOKEN.\n")
	}
	fmt.Printf("Token => '%v'\n", token)
	return
}

func GetFormattedToken() string {
	return fmt.Sprintf("Bot %v", getBotToken())
}

func initBot() (err error) {
	fmt.Printf("Digona (version: %v), initialization...\n", digonaVersion)
	DigonaBot.session, err = discordgo.New(GetFormattedToken())
	if err != nil {
		log.Fatalf("An error occured at the bot creation: %v\n", err.Error())
	}
	DigonaBot.data, err = DigonaBot.session.User("@me")
	if err != nil {
		log.Printf("Cannot retrieve own user's info: %v\n", err.Error())
	}
	DigonaBot.upTime = time.Now().UTC()
}