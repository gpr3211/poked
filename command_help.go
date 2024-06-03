package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func commandHelp(c *config, s ...string) error {

	HeaderStyle := lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Center)
	EvenRowStyle := lipgloss.NewStyle().Padding(0, 1)
	OddRowStyle := lipgloss.NewStyle().Padding(0, 1)
	rows := [][]string{}

	for _, cmd := range getCommands() {
		rows = append(rows, []string{cmd.name, cmd.description})
	}
	t := table.New().Border(lipgloss.NormalBorder()).BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers("Command", "Description").
		Rows(rows...)
	fmt.Println(t)

	return nil
}
