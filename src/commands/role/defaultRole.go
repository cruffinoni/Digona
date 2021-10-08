package role

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"time"
)

func getAllMembersWithoutRole(guildId string) ([]*discordgo.Member, error) {
	members, err := skeleton.Bot.GetSession().GuildMembers(guildId, "", 1000)
	if err != nil {
		return nil, err
	}
	users := make([]*discordgo.Member, 0)
	for _, m := range members {
		if len(m.Roles) == 0 {
			users = append(users, m)
		}
	}
	return users, nil
}

func attributeRoleToMembers(role *discordgo.Role, members []*discordgo.Member, guildId string) error {
	skeleton.Bot.Logf("There is %v member to attribute a role\n", len(members))
	for k, i := range members {
		if err := skeleton.Bot.GetSession().GuildMemberRoleAdd(guildId, i.User.ID, role.ID); err != nil {
			skeleton.Bot.Errorf("An error occurred while setting the role '%v' to '%v': %v\n", role.Mention(), i.User.ID, err)
			return err
		}
		if k%50 == 0 {
			time.Sleep(10 * time.Second)
		}
	}
	return nil
}

func SetDefaultRole(parser *parser.MessageParser) error {
	args := parser.GetArguments()
	if len(args) != 1 {
		skeleton.Bot.SendDelayedMessage(parser.GetChannelId(), "Entrez uniquement le rôle à attribuer")
		return nil
	}
	var role *discordgo.Role = nil
	roles, err := skeleton.Bot.GetSession().GuildRoles(parser.GetGuildId())
	if err != nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}
	if role = discord.FindRoleFromRawRoleId(roles, args[0]); role == nil {
		skeleton.Bot.SendMessage(parser.GetChannelId(), fmt.Sprintf("Impossible de trouver le rôle: '%v'", args[0]))
		return err
	}
	members, err := getAllMembersWithoutRole(parser.GetGuildId())
	if err != nil {
		return err
	}
	if len(members) == 0 {
		skeleton.Bot.SendDelayedMessage(parser.GetChannelId(), "Aucun utilisateur valide n'a été trouvé")
		return discord.DeleteMessage(parser.GetChannelId(), parser.GetDiscordMessage().ID)
	}
	if err = attributeRoleToMembers(role, members, parser.GetGuildId()); err != nil {
		skeleton.Bot.SendInternalServerErrorMessage(parser.GetGuildId())
		return err
	}
	skeleton.Bot.SendMessage(parser.GetChannelId(), fmt.Sprintf("J'ai ajouté le rôle %v à %v personne(s)", role.Mention(), len(members)))
	return discord.DeleteMessage(parser.GetChannelId(), parser.GetDiscordMessage().ID)
}
