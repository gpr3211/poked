package main

import (
	"fmt"
	"log"
	"main/pokeapi"
)

func callMap(cfg *config) error {
	pokeapiClient := pokeapi.NewClient()
	resp, err := pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		log.Fatal(resp)
	}
	fmt.Println("Location  areas list:")
	for _, loc := range resp.Results {
		fmt.Printf("%s, \n", loc.Name)
	}
	cfg.nextLocationURL = resp.Next

	cfg.prevLocationURL = resp.Previous
	return nil
}
