package commands

import (
	"fmt"
	"github.com/Digona/leagueOfLegends"
	//"github.com/Digona/digona"
)

func ShowMostPlayedChamp(parser *MessageParser) error {
	//var lastError error
	playersList := make([]leagueOfLegends.PlayerStructure, leagueOfLegends.MaxPlayerPerGame)
	for _, playerName := range parser.GetArguments() {
		if currentPlayer, err := leagueOfLegends.GetPlayerData(playerName); err == nil {
			playersList = append(playersList, currentPlayer)
			err = currentPlayer.RetrieveLastPlayedChamp()
			if err != nil {
				fmt.Printf("Unable to retrieve last played champ, error: %v\n", err.Error())
			}
			fmt.Printf("Player %v (IDs: '%v' - '%v') added to the list\n", currentPlayer.Name, currentPlayer.Id, currentPlayer.AccountId)
		} else {
			fmt.Printf("An error occured during retrieved data of %v -> %v\n", playerName, err.Error())
			continue
		}
	}
	return nil
}
