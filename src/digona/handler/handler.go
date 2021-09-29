package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/hooks"
)

var discordHandlers = []interface{}{
	onBotReady,
	hooks.OnMessageCreated,
	hooks.OnUserReact,
	hooks.OnUserRemoveReact,
}

func RegisterHandler(session *discordgo.Session) {
	for _, function := range discordHandlers {
		session.AddHandler(function)
	}
}

func onBotReady(_ *discordgo.Session, _ *discordgo.Ready) {
	skeleton.Bot.Log("Digona's discord session is ready!\n")
}
