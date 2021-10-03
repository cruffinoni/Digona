package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

func OnBotReady(session *discordgo.Session, ready *discordgo.Ready) {
	for _, guild := range ready.Guilds {
		fullGuild, err := session.Guild(guild.ID)
		if err != nil {
			skeleton.Bot.Fatalf("unable to get guild's data for id %v\n", fullGuild.ID)
		}
		skeleton.Bot.RegisterGuild(fullGuild)
		skeleton.Bot.Logf("Guild '%v' (id: %v) registered\n", fullGuild.Name, fullGuild.ID)
	}
	skeleton.Bot.Log("Digona's discord session is ready!\n")
}
