package main

import (
	"errors"
	"fmt"
)

func callInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No names were provided")
	}
	name := args[0]
	c, ok := cfg.pokemonCaught[name]
	if !ok {
		return errors.New("you haven't caught the pokemon yet")
	}
	stats := c.Stats
	fmt.Printf("Name: %s\n", c.Species.Name)
	fmt.Printf("Height: %v\n", c.Height)
	fmt.Printf("Weight: %v\n", c.Weight)
	for _, stat := range stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range c.Types {
		fmt.Println(" - ", typ.Type.Name)
	}

	fmt.Println("")
	return nil
}
