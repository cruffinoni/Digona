package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/reaction"
	"github.com/cruffinoni/Digona/src/digona/config"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
)

func OnBotReady(_ *discordgo.Session, _ *discordgo.Ready) {
	skeleton.Bot.Log("Digona's discord session is ready!\n")
}

func OnGuildCreate(_ *discordgo.Session, guild *discordgo.GuildCreate) {
	skeleton.Bot.RegisterGuild(guild.Guild)
	skeleton.Bot.Logf("Guild '%v' added\n", guild.Name)
	if !config.FileExists(guild.ID) {
		if err := config.Create(guild.ID); err != nil {
			skeleton.Bot.Errorf("unable to create a config file (guild id %v) => \n", guild.ID, err)
		}
	}
	if err := reaction.LoadReactionMessage(guild.ID); err != nil {
		skeleton.Bot.Errorf("unable to load the reaction message (guild id %v): %v\n", guild.ID, err)
	}
}

func OnGuildDelete(_ *discordgo.Session, guild *discordgo.GuildDelete) {
	skeleton.Bot.RemoveGuild(guild.ID)
	skeleton.Bot.Logf("Guild id %v deleted\n", guild.ID)
}
