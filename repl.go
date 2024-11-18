package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	cfg := &config{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pokedex > ")
	fmt.Println("Type 'help' to see available commands.")
	fmt.Println("Type 'exit' to quit.")

	initCommands()

	for {
		fmt.Print("pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if cmd, exists := commands[input]; exists {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next set of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous set of locations",
			callback:    commandMapB,
		},
	}
}
