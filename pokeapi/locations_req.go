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
	data, ok := c.cache.Get(full)
	if ok { //cache hit
		fmt.Println("cache hit")
		locations := LocationsList{}

		err := json.Unmarshal(data, &locations)
		fmt.Println("json done")
		if err != nil {
			return LocationsList{}, err
		}
		return locations, nil
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
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationsList{}, err
	}
	locations := LocationsList{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationsList{}, err
	}
	c.cache.Add(full, data)
	return locations, nil
}
func (c *Client) ListArea(LocationAreaName string) (LocationArea, error) { // Method on client

	endpoint := "/location-area/" + LocationAreaName
	full := baseURL + endpoint

	data, ok := c.cache.Get(full)
	if ok { //cache hit
		fmt.Println("cache hit")
		locationAr := LocationArea{}

		err := json.Unmarshal(data, &locationAr)
		fmt.Println("json done")
		if err != nil {
			return LocationArea{}, err
		}
		return locationAr, nil
	}

	req, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(full, data)
	return locations, nil
}
func (c *Client) ListStats(pokename string) (PokemonInfo, error) { // Method on client

	endpoint := "/pokemon/" + pokename
	full := baseURL + endpoint

	data, ok := c.cache.Get(full)
	if ok { //cache hit
		fmt.Println("cache hit")
		locationAr := PokemonInfo{}

		err := json.Unmarshal(data, &locationAr)
		fmt.Println("json done")
		if err != nil {
			return PokemonInfo{}, err
		}
		return locationAr, nil
	}

	req, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return PokemonInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	if resp.StatusCode > 399 {
		return PokemonInfo{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	locations := PokemonInfo{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return PokemonInfo{}, err
	}
	c.cache.Add(full, data)
	return locations, nil
}
