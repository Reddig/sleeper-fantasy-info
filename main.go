package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sleeper-fantasy-info/internal/clients/sleeperapi"
	"sleeper-fantasy-info/models"
)

type LeaguesJSON struct {
	Leagues []string `json:"leagues"`
}

func main() {
	client := sleeperapi.NewClient("https://api.sleeper.app/v1/")
	data, err := os.ReadFile("leagues.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var leaguesJSON LeaguesJSON
	err = json.Unmarshal(data, &leaguesJSON)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	allLeagues := make([]models.League, len(leaguesJSON.Leagues))
	var cnt int = 0
	for _, leagueID := range leaguesJSON.Leagues {
		league, err := client.GetLeague(leagueID)
		if err != nil {
			log.Fatalf("Error getting league with ID %s: %v", leagueID, err)
		}
		allLeagues[cnt] = *league
		fmt.Printf("League: %s (%s)\n", allLeagues[cnt].Name, allLeagues[cnt].Season)
		cnt++
	}

}
