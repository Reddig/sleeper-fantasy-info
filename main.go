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
var leagueStorage = map[string]*models.League{}
var playerStorage = map[string]*models.Player{}
var userStorage = map[string]*models.User{}
var rosterStorage = map[int]*models.Roster{}

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
	for _, leagueID := range leaguesJSON.Leagues {
		league, err := client.GetLeague(leagueID)
		if err != nil {
			log.Fatalf("Error getting league with ID %s: %v", leagueID, err)
		}
		leagueStorage[leagueID] = league
		fmt.Printf("League: %s (%s)\n", leagueStorage[leagueID].Name, leagueStorage[leagueID].Season)
	}
	
	players, err := client.GetPlayers()
	if err != nil {
		log.Fatalf("Error getting players: %v", err)
	}
	for _, player := range players {
		playerStorage[player.PlayerID] = &player
	}
	for k := range(playerStorage) {
		fmt.Printf("A Player: %s %s (%s)\n", playerStorage[k].FirstName, playerStorage[k].LastName, playerStorage[k].PlayerID)
		break
	}
	for _, league := range leagueStorage {
		rosters, err := client.GetRosters(league.LeagueID)
		if err != nil {
			log.Fatalf("Error getting rosters: %v", err)
		}
		for _, roster := range rosters {
			rosterStorage[roster.RosterID] = &roster
		}
	}

	for _, roster := range rosterStorage {
		user, err := client.GetUser(roster.OwnerID)
		if err != nil {
			log.Fatalf("Error getting user with ID %s: %v", roster.OwnerID, err)
		}
		userStorage[user.UserID] = user
		fmt.Printf("A user: %s (%s)\n", user.DisplayName, user.UserID)
	}

	

	for _, league := range leagueStorage {
		for i := 1; i <= 18; i++ {
			transactions, err := client.GetTransactionsForWeek(league.LeagueID, i)
			if err != nil {
				log.Fatalf("Error getting transactions for %s in week %d: %v", league.Name, i, err)
			}
			for _, t := range transactions {
				if len(t.Adds) > 0 {
					for k := range t.Adds {
						var roster = rosterStorage[t.Adds[k]]
						var userID = roster.OwnerID
						if err != nil {
							log.Fatalf("Error getting user with ID %s: %v", userID, err)
						}
						// fmt.Printf("%s week %d %s %s acquired by %s\n", league.Season, t.Leg, playerStorage[k].FirstName, playerStorage[k].LastName, userStorage[userID].DisplayName)
						playerStorage[k].Transactions[t.TransactionID] = t
					}
				}
			}
		}
	}

	var mostTradedPlayers = make(map[*models.Player]int)
	var mostTradedCount = 0
	var mostTransactedPlayers = make(map[*models.Player]int)
	var mostTransactedCount = 0
	for _, player := range playerStorage {
		if len(player.Transactions) > mostTransactedCount{
			fmt.Printf("%s %s was involved in %d transactions, higher than current max of %d\n", player.FirstName, player.LastName, len(player.Transactions), mostTransactedCount)
			mostTransactedCount = len(player.Transactions)
			mostTransactedPlayers =  make(map[*models.Player]int)
			mostTransactedPlayers[player] = len(player.Transactions)
		} else if len(player.Transactions) == mostTransactedCount {
			fmt.Printf("%s %s was involved in %d transactions, matching the current max of %d\n", player.FirstName, player.LastName, len(player.Transactions), mostTransactedCount)
			mostTransactedPlayers[player] = len(player.Transactions)
		}
		var tradeCount = 0
		for _, t := range player.Transactions {
			if t.Type == "trade" {
				tradeCount++
			}
		}
		if tradeCount > mostTradedCount {
			fmt.Printf("%s %s was involved in %d trades, higher than current max of %d\n", player.FirstName, player.LastName, tradeCount, mostTradedCount)
			mostTradedCount = tradeCount
			mostTradedPlayers =  make(map[*models.Player]int)
			mostTradedPlayers[player] = tradeCount
		} else if tradeCount == mostTradedCount {
			fmt.Printf("%s %s was involved in %d trades, matching the current max of %d\n", player.FirstName, player.LastName, tradeCount, mostTradedCount)
			mostTradedPlayers[player] = tradeCount
		}
	}
	fmt.Printf("\nThe highest number of trades for a single player is %d\nThe following players have been traded %d times.\n", mostTradedCount, mostTradedCount)
	for k := range mostTradedPlayers{
		fmt.Printf("%s %s\n", k.FirstName, k.LastName)
	}
	fmt.Printf("\nThe highest number of moves for a single player is %d\nThe following players have been moved %d times.\n", mostTransactedCount, mostTransactedCount)
	for k := range mostTransactedPlayers{
		fmt.Printf("%s %s\n", k.FirstName, k.LastName)
	}
}