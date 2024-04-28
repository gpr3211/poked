package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	name := args[0]

	resp, err := cfg.pokeapiClient.ListStats(name)
	base := resp.BaseExperience
	randomseed := rand.Intn(base * 2)
	if base > randomseed {
		fmt.Printf("You failed to catch %s\n", resp.Name)
	} else {
		fmt.Printf("You caught %s!\n", resp.Name)
	}
	if err != nil {
		log.Fatal(resp)
	}
	fmt.Println("")

	fmt.Println("")
	fmt.Println("")

	return nil

}