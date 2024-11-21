package sleeperapi

import (
	"encoding/json"
	// "errors"
	"fmt"
	// "io/ioutil"
	// "net/http"
	"sleeper-fantasy-info/models"
)

func (c *SleeperClient) GetLeague(leagueID string) (*models.League, error) {
	// Construct the URL
	endpoint := fmt.Sprintf("/league/%s", leagueID)
	resp, err := c.MakeRequest(endpoint, "GET", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// Parse the response body
	var League models.League
	if err := json.NewDecoder(resp.Body).Decode(&League); err != nil {
		return nil, err
	}

	return &League, nil
}
