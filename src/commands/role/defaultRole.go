package role

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"github.com/cruffinoni/Digona/src/discord"
	"regexp"
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
			time.Sleep(3)
		}
	}
	return nil
}

func SetDefaultRole(msg *parser.MessageParser) error {
	var role *discordgo.Role = nil
	roles, err := skeleton.Bot.GetSession().GuildRoles(msg.GetGuildId())
	if err != nil {
		skeleton.Bot.SendMessage(msg.GetChannelId(), "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}
	for _, i := range msg.GetArguments() {
		if matched, err := regexp.Match("<@&\\d{18}>", []byte(i)); err != nil {
			return err
		} else if matched {
			if role = discord.FindRoleFromRawRoleId(roles, i); role == nil {
				skeleton.Bot.SendMessage(msg.GetChannelId(), fmt.Sprintf("Impossible de trouver le rôle: '%v'", i))
				return err
			}
			break
		}
	}
	if role == nil {
		skeleton.Bot.SendDelayedMessage(msg.GetChannelId(), "Aucun role n'a été trouvé dans le message.")
		return nil
	}
	members, err := getAllMembersWithoutRole(msg.GetGuildId())
	if err != nil {
		return err
	}
	if len(members) == 0 {
		skeleton.Bot.SendDelayedMessage(msg.GetChannelId(), "Aucun utilisateur valide n'a été trouvé")
		return discord.DeleteMessage(msg.GetChannelId(), msg.GetDiscordMessage().ID)
	}
	if err = attributeRoleToMembers(role, members, msg.GetGuildId()); err != nil {
		skeleton.Bot.SendInternalServerErrorMessage(msg.GetGuildId())
		return err
	}
	skeleton.Bot.SendMessage(msg.GetChannelId(), fmt.Sprintf("J'ai ajouté le rôle %v à %v personne(s)", role.Mention(), len(members)))
	return discord.DeleteMessage(msg.GetChannelId(), msg.GetDiscordMessage().ID)
}
