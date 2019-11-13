package leagueOfLegends

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MatchFormat struct {
	Lane      string `json:"lane"`
	GameId    uint64 `json:"gameId"`
	Champion  int    `json:"champion"`
	Region    string `json:"PlatformId"`
	TimeStamp uint64 `json:"timestamp"`
	Queue     uint32 `json:"queue"`
	Role      string `json:"role"`
	Season    uint   `json:"season"`
}

type ResponseMapMatch struct {
	Matches    []MatchFormat `json:"matches"`
	EndIndex   uint32        `json:"endIndex"`
	StartIndex uint32        `json:"startIndex"`
	TotalGames uint32        `json:"totalGames"`
}

const (
	apiEndpointLastMatch = "/lol/match/v4/matchlists/by-account/%v?season=13&%v"
)

func (player PlayerStructure) RetrieveLastPlayedChamp() error {
	res, err := http.Get(fmt.Sprintf(apiRiotEuwRegion+apiEndpointLastMatch, player.AccountId, apiRiotKey))
	fmt.Print(fmt.Sprintf(apiRiotEuwRegion+apiEndpointLastMatch, player.AccountId, apiRiotKey))
	if err != nil || res.StatusCode != http.StatusOK {
		if res != nil && res.StatusCode != http.StatusOK {
			log.Printf("Status code should be %v but is %v instead\n", http.StatusOK, res.StatusCode)
			return errors.New(http.StatusText(res.StatusCode))
		}
		return err
	}
	var playerData []byte
	playerData, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Unable to read the body request for %v\n", player.Name)
		return err
	}
	listAllChampions, err := RetrieveAllChampions()
	playerPlayedChampion := make(map[int]*PlayerPlayedChamp)
	if err != nil {
		return nil
	}
	var playerMatches ResponseMapMatch
	if err = json.Unmarshal(playerData, &playerMatches); err != nil {
		log.Printf("Unable to unmarshal data %v\n", player.Name)
		return err
	}
	for _, match := range playerMatches.Matches {
		if playerPlayedChampion[match.Champion] == nil {
			playerPlayedChampion[match.Champion] = &PlayerPlayedChamp{
				Lane:         nil,
				Role:         nil,
				Occurrence:   0,
				ChampionName: "none",
				Win:          0,
				Lose:         0,
			}
		}
		playerPlayedChampion[match.Champion].Occurrence += 1
		playerPlayedChampion[match.Champion].Lane = append(playerPlayedChampion[match.Champion].Lane, match.Lane)
		playerPlayedChampion[match.Champion].Role = append(playerPlayedChampion[match.Champion].Role, match.Role)
		fmt.Printf("[%v] Played champion id %v (%v) on %v (as %v)\n", match.GameId, match.Champion,
			listAllChampions[match.Champion], match.Lane, match.Role)
	}
	for i := range playerPlayedChampion {
		//fmt.Printf("i: %v & id: %v\n", i, championId)
		fmt.Printf("Played %v %v times on %+v\n", listAllChampions[i],
			playerPlayedChampion[i].Occurrence, playerPlayedChampion[i].GetAllLaneName())
	}
	return nil
}
