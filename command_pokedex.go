package main

import (
	"fmt"
)

func callPokedex(cfg *config, args ...string) error {
	list := cfg.pokemonCaught
	fmt.Println("List of pokemons in bag :")
	for _, item := range list {
		fmt.Printf(" - %s\n", item.Species.Name)
	}
	fmt.Println("")
	return nil
}
