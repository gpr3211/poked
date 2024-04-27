package main

import (
	"fmt"
	"log"
)

func callMapb(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		log.Fatal(resp)
	}
	fmt.Println("Location Areas: ")
	fmt.Println()
	for _, loc := range resp.Results {
		fmt.Printf("%s, \n", loc.Name)
	}
	fmt.Println("")
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil

}
