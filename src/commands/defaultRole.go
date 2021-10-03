package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cruffinoni/Digona/src/digona/skeleton"
	"regexp"
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
	fmt.Printf("There is %v member to attribute a role\n", len(members))
	for _, i := range members {
		if err := skeleton.Bot.GetSession().GuildMemberRoleAdd(guildId, i.User.ID, role.ID); err != nil {
			skeleton.Bot.Errorf("An error occurred while setting the role '%v' to '%v': %v\n", role.Mention(), i.User.ID, err)
			return err
		}
	}
	return nil
}

func SetDefaultRole(msg *MessageParser) error {
	var role *discordgo.Role = nil
	roles, err := skeleton.Bot.GetSession().GuildRoles(msg.guildId)
	if err != nil {
		skeleton.Bot.SendMessage(msg.channel, "Je ne peux pas récupérer les roles de ce serveur")
		return err
	}
	for _, i := range msg.args {
		if matched, err := regexp.Match("<@&\\d{18}>", []byte(i)); err != nil {
			return err
		} else if matched {
			if role = retrieveRole(roles, i); role == nil {
				skeleton.Bot.SendMessage(msg.channel, fmt.Sprintf("Impossible de trouver le rôle: '%v'", i))
				return err
			}
			break
		}
	}
	if role == nil {
		skeleton.Bot.SendDelayedMessage(msg.channel, "Aucun role n'a été trouvé dans le message.")
		return nil
	}
	members, err := getAllMembersWithoutRole(msg.guildId)
	if err != nil {
		return err
	}
	if len(members) == 0 {
		skeleton.Bot.SendMessage(msg.channel, "Aucun utilisateur valide n'a été trouvé")
		return deleteLastMessage(msg.channel, msg.message.ID)
	}
	if err = attributeRoleToMembers(role, members, msg.guildId); err != nil {
		skeleton.Bot.SendInternalServerErrorMessage(msg.guildId)
		return err
	}
	skeleton.Bot.SendMessage(msg.channel, fmt.Sprintf("J'ai ajouté le rôle %v à %v personne(s)", role.Mention(), len(members)))
	return deleteLastMessage(msg.channel, msg.message.ID)
}
