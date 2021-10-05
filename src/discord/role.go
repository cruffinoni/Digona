package discord

import "github.com/bwmarrin/discordgo"

func GetRoleFromRawRoleId(roleId string) string {
	return roleId[3 : len(roleId)-1]
}

func FindRoleFromRawRoleId(roles []*discordgo.Role, reference string) *discordgo.Role {
	reference = GetRoleFromRawRoleId(reference)
	for _, r := range roles {
		if r.ID == reference {
			return r
		}
	}
	return nil
}
