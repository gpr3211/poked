package main

import (
	"fmt"
)

func commandHelp(c *config, s ...string) error {
	fmt.Println("List of commands")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s:   %s\n", cmd.name, cmd.description)
		fmt.Println()
	}
	return nil
}
