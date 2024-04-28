package main

import (
	"main/pokeapi"
	"time"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string // pointers bc we can haf nil
	prevLocationURL *string
	pokemonCaught   map[string]pokeapi.PokemonInfo
}

func main() {
	// init client. ommiting prev and onext at init sets them to nil
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	cfg := config{
		pokeapiClient: pokeClient,
		pokemonCaught: make(map[string]pokeapi.PokemonInfo),
	}

	startRep(&cfg)
}
