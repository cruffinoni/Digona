package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/hooks"
)

var discordHandlers = []interface{}{
	hooks.OnBotReady,
	hooks.OnMessageCreated,
	hooks.OnUserReact,
	hooks.OnUserRemoveReact,
	hooks.OnGuildCreate,
	hooks.OnGuildDelete,
	hooks.OnUserUseSlashCommands,
}

func RegisterHandler(session *discordgo.Session) {
	for _, function := range discordHandlers {
		session.AddHandler(function)
	}
}
