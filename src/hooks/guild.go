package hooks

import (
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/reaction"
	"github.com/cruffinoni/Digona/src/digona/config"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"github.com/cruffinoni/Digona/src/user"
)

func OnBotReady(_ *discordgo.Session, _ *discordgo.Ready) {
	skeleton.Bot.Log("Digona's discord session is ready!\n")
}

func OnGuildCreate(_ *discordgo.Session, guild *discordgo.GuildCreate) {
	skeleton.Bot.RegisterGuild(guild.Guild)
	skeleton.Bot.Logf("Guild '%v' added\n", guild.Name)
	usersFromGuild, err := skeleton.Bot.GetDatabase().LoadUsersForGuild(guild.ID)
	if err != nil {
		skeleton.Bot.Errorf("unable to get all users from the database (guild id %v): %v\n", guild.ID, err)
		return
	}
	user.StoreUsersFromModels(guild.ID, usersFromGuild)
	var users []*discordgo.Member
	if users, err = discord.GetAllUsersFromGuild(guild.ID); err != nil {
		skeleton.Bot.Errorf("unable to get all users from the guild (guild id %v): %v\n", guild.ID, err)
	} else {
		for _, u := range users {
			if _, err = skeleton.Bot.GetDatabase().AddUser(guild.ID, u.User.ID); err != nil {
				skeleton.Bot.Errorf("can't user to database (user id %v / guild id %v): %v\n", u.User.ID, guild.ID, err)
			}
		}
	}
	if !config.FileExists(guild.ID) {
		if err := config.Create(guild.ID); err != nil {
			skeleton.Bot.Errorf("unable to create a config file (guild id %v) => \n", guild.ID, err)
		}
	}
	if err = reaction.LoadReactionMessage(guild.ID); err != nil {
		skeleton.Bot.Errorf("unable to load the reaction message (guild id %v): %v\n", guild.ID, err)
	}
}

func OnGuildDelete(_ *discordgo.Session, guild *discordgo.GuildDelete) {
	skeleton.Bot.RemoveGuild(guild.ID)
	skeleton.Bot.Logf("Guild id %v deleted\n", guild.ID)
	if err := skeleton.Bot.GetDatabase().DeleteUsersFromGuild(guild.ID); err != nil {
		skeleton.Bot.Errorf("can't delete all users from database (guild id %v): %v\n", guild.ID, err)
	}
}
