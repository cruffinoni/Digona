package leagueOfLegends

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PlayerStructure struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId uint64 `json:"profileIconId"`
	RevisionDate  uint64 `json:"revisionDate"`
	SummonerLevel uint64 `json:"summonerLevel"`
}

const (
	MaxPlayerPerGame = 10

	apiEndpointSummoners = "/lol/summoner/v4/summoners/by-name/%v?%v"
	apiRiotEuwRegion     = "https://euw1.api.riotgames.com"
	apiRiotKey           = "api_key=RGAPI-b0624a8a-2afa-4415-98cc-d72fb2b9e864"
)

func GetPlayerData(name string) (PlayerStructure, error) {
	res, err := http.Get(fmt.Sprintf(apiRiotEuwRegion+apiEndpointSummoners, name, apiRiotKey))
	if err != nil || res.StatusCode != http.StatusOK {
		if res != nil && res.StatusCode != http.StatusOK {
			log.Printf("Status code should be %v but is %v instead\n", http.StatusOK, res.StatusCode)
			return PlayerStructure{}, errors.New(http.StatusText(res.StatusCode))
		}
		return PlayerStructure{}, err
	}
	var playerData []byte
	playerData, err = ioutil.ReadAll(res.Body)
	//fmt.Printf("Body from the response: '%v'\n", string(playerData))
	if err != nil {
		log.Printf("Unable to read the body request for %v\n", name)
		return PlayerStructure{}, err
	}
	var currentPlayer PlayerStructure
	if err = json.Unmarshal(playerData, &currentPlayer); err != nil {
		log.Printf("Unable to unmarshal data %v\n", name)
		return PlayerStructure{}, err
	}
	return currentPlayer, nil
}
