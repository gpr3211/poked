package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRep(cfg *config) error {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, type 'help' to see a list of commands ")
	for {
		fmt.Println("")
		fmt.Print("Enter command: ")
		fmt.Println("")
		reader.Scan()
		words := normInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		commands, exists := getCommands()[command]
		if exists {
			err := commands.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func normInput(text string) []string {
	low := strings.ToLower(text)
	sep := strings.Fields(low)
	return sep
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display locations/page forward",
			callback:    callMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display locations/page backward",
			callback:    callMapb,
		},
		"explore": {
			name:        "explore",
			description: "show all pokemon that populate an area",
			callback:    callExplore,
		},
		"catch": {
			name:        "explore",
			description: "type catch {pokemonname} to attemp to catch",
			callback:    callCatch,
		},
	}
}
