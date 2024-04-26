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
		fmt.Print("Enter command: ")
		reader.Scan()
		words := normInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		commands, exists := getCommands()[command]
		if exists {
			err := commands.callback(cfg)
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
	id          int
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			id:          0,
			name:        "help",
			description: "Displays help",
			callback:    commandHelp,
		},
		"exit": {
			id:          1,
			name:        "exit",
			description: "Exit program",
			callback:    commandExit,
		},
		"map": {
			id:          2,
			name:        "map",
			description: "Display locations/page forward",
			callback:    callMap,
		},
		"mapb": {
			id:          2,
			name:        "mapb",
			description: "Display locations/page backward",
			callback:    callMapb,
		},
	}

}
