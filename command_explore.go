package main

import (
	"errors"
	"fmt"
	"log"
)

func callExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.ListArea(locationAreaName)
	if err != nil {
		log.Fatal(resp)
	}
	fmt.Println("")
	fmt.Printf("Pokemon in %s : ", locationAreaName)
	fmt.Println("")
	for i, loc := range resp.PokemonEncounters {
		fmt.Printf("%v - %s \n", i+1, loc.Pokemon.Name)
	}
	fmt.Println("")

	return nil

}
