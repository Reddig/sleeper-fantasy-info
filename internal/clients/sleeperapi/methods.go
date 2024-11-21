package sleeperapi

import (
	"encoding/json"
	// "log"
	"net/http"

	// "errors"
	"fmt"
	// "io/ioutil"
	// "net/http"
	"sleeper-fantasy-info/models"
)

func Get(client *SleeperClient, endpoint string) (*http.Response, error) {
	resp, err := client.MakeRequest(endpoint, "GET", nil)
	if err != nil {
		return nil, err
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}
	return resp, nil
}

func ParseResponse[T any](resp *http.Response, respStruct *T) (error) {
	defer resp.Body.Close()
	// attempt to parse the response to the struct given, else fail
	if err := json.NewDecoder(resp.Body).Decode(&respStruct); err != nil {
		return err
	}
	return nil
}

func (c *SleeperClient) GetLeague(leagueID string) (*models.League, error) {
	resp, err := Get(c, fmt.Sprintf("/league/%s", leagueID))
	if err != nil {
		return nil, err
	} 
	var league models.League
	if err := ParseResponse(resp, &league); err != nil {
		return nil, err
	}

	return &league, nil
}

func (c *SleeperClient) GetTransactionsForWeek(leagueID string, week int) ([]models.Transaction, error) {
	resp, err := Get(c, fmt.Sprintf("/league/%s/transactions/%d", leagueID, week))
	if err != nil {
		return nil, err
	} 
	var transactions []models.Transaction
	if err := ParseResponse(resp, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (c *SleeperClient) GetPlayers() (map[string]models.Player, error) {
	resp, err := Get(c, fmt.Sprintf("/players/nfl"))
	if err != nil {
		return nil, err
	} 
	var players map[string]models.Player
	if err := ParseResponse(resp, &players); err != nil {
		return nil, err
	}

	return players, nil
}