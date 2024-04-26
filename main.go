package main

import "main/pokeapi"

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string // pointers bc we can haf nil
	prevLocationURL *string
}

func main() {
	// init client. ommiting prev and next at init sets them to nil
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRep(&cfg)
}
