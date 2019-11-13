package leagueOfLegends

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ImageData struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}
type ChampionDataInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}
type ChampionDataStats struct {
	HealthPoints                    float64 `json:"hp"`
	HealthPointsPerLevel            float64 `json:"hpperlevel"`
	ManaPoints                      float64 `json:"mp"`
	ManaPointsPerLevel              float64 `json:"mpperlevel"`
	MovementSpeed                   float64 `json:"movespeed"`
	Armor                           float64 `json:"armor"`
	ArmorPerLevel                   float64 `json:"armorperlevel"`
	SpellBlock                      float64 `json:"spellblock"`
	SpellBlockPerLevel              float64 `json:"spellblockperlevel"`
	AttackRange                     float64 `json:"attackrange"`
	HealthPointRegeneration         float64 `json:"hpregen"`
	HealthPointRegenerationPerLevel float64 `json:"hpregenperlevel"`
	ManaPointRegeneration           float64 `json:"mpregen"`
	ManaPointRegenerationPerLevel   float64 `json:"mpregenperlevel"`
	CriticalStrikeChance            float64 `json:"crit"`
	CriticalStrikeChancePerLevel    float64 `json:"critperlevel"`
	AttackDamage                    float64 `json:"attackdamage"`
	AttackDamagePerLevel            float64 `json:"attackdamageperlevel"`
	AttackSpeedPerLevel             float64 `json:"attackspeedperlevel"`
	AttackSpeed                     float64 `json:"attackspeed"`
}

type ChampionData struct {
	Version string            `json:"version"`
	ID      string            `json:"id"`
	Key     string            `json:"key"`
	Name    string            `json:"name"`
	Title   string            `json:"title"`
	Blurb   string            `json:"blurb"`
	Info    ChampionDataInfo  `json:"info"`
	Image   ImageData         `json:"image"`
	Tags    []string          `json:"tags"`
	Partype string            `json:"partype"`
	Stats   ChampionDataStats `json:"stats"`
}

type ChampionRequest struct {
	FileType string                  `json:"type"`
	Format   string                  `json:"format"`
	Version  string                  `json:"version"`
	Data     map[string]ChampionData `json:"data"`
}

type ChampionList map[int]string

type PlayerPlayedChamp struct {
	Lane       []string
	Role       []string
	Occurrence int
	ChampionName string
	Win uint
	Lose uint
}

const (
	apiEndpointChampionData = "http://ddragon.leagueoflegends.com/cdn/9.18.1/data/en_US/champion.json"
)

func RetrieveAllChampions() (ChampionList, error) {
	res, err := http.Get(apiEndpointChampionData)
	if err != nil || res.StatusCode != http.StatusOK {
		if res != nil && res.StatusCode != http.StatusOK {
			log.Printf("Status code should be %v but is %v instead\n", http.StatusOK, res.StatusCode)
			return ChampionList{}, errors.New(http.StatusText(res.StatusCode))
		}
		return ChampionList{}, err
	}
	var requestBody []byte
	requestBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Unable to read the body request\n")
		return ChampionList{}, err
	}
	var requestResult ChampionRequest
	if err = json.Unmarshal(requestBody, &requestResult); err != nil {
		log.Printf("Unable to unmarshal data\n")
		return ChampionList{}, err
	}
	championList := make(ChampionList)
	for _, champion := range requestResult.Data {
		key, err := strconv.Atoi(champion.Key)
		if err != nil {
			log.Printf("Unable to convert str to int\n")
			continue
		}
		championList[key] = champion.Name
		//fmt.Printf("(id %v) %v\n", champion.Key, champion.Name)
	}
	return championList, nil
}

func (ppc PlayerPlayedChamp) GetAllLaneName() []string {
	lanesName := make([]string, len(ppc.Lane))
	for _, rawLaneName := range ppc.Lane {
		switch rawLaneName {
		case "TOP":
			lanesName = append(lanesName, "Top")
		case "MID":
			lanesName = append(lanesName, "Mid")
		case "JUNGLE":
			lanesName = append(lanesName, "Jungle")
		case "BOTTOM":
			lanesName = append(lanesName, "Bot")
		case "NONE":
			lanesName = append(lanesName, "Aucune")
		}
	}
	return lanesName
}
