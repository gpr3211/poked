package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	name := args[0]
	resp, err := cfg.pokeapiClient.ListStats(name)
	if err != nil {
		return errors.New("error")
	}
	base := resp.BaseExperience
	fmt.Printf("Throwing a pokeball at %s ...\n", resp.Name)
	randomseed := rand.Intn(base * 2)
	if base > randomseed {
		fmt.Printf("You failed to catch %s\n", resp.Name)
	} else {
		fmt.Printf("You caught %s !!!\n", resp.Name)
		cfg.pokemonCaught[resp.Name] = resp
	}

	fmt.Println("")
	fmt.Println("")
	return nil
}
