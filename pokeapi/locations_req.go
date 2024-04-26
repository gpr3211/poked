package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(page *string) (LocationsList, error) { // Method on client

	endpoint := "/location-area"
	full := baseURL + endpoint

	if page != nil {
		full = *page
	}
	req, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return LocationsList{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsList{}, err
	}
	if resp.StatusCode > 399 {
		return LocationsList{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsList{}, err
	}
	locations := LocationsList{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsList{}, err
	}
	return locations, nil

}
