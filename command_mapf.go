package main

import (
	"fmt"
	"log"
)

func callMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		log.Fatal(resp)
	}

	fmt.Println("Location  areas list:")
	fmt.Println("")
	for _, loc := range resp.Results {
		fmt.Printf("%s, \n", loc.Name)
	}
	cfg.nextLocationURL = resp.Next
	fmt.Println("")
	cfg.prevLocationURL = resp.Previous
	return nil
}

func callMapb(cfg *config, args ...string) error {
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
