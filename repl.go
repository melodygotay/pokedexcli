package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func startRepl(cfg *config) {
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

		// Split input into command and arguments
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		commandName := parts[0]
		args := parts[1:] // This handles additional arguments, such as area names

		if cmd, exists := commands[commandName]; exists {
			if err := cmd.callback(cfg, args...); err != nil {
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
	callback    func(*config, ...string) error
}

func commandCatchWrapper(cfg *config, args ...string) error {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return commandCatch(cfg, rnd, args...)
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
		"explore": {
			name:        "explore",
			description: "Displays the Pokemon in a specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the specified Pokemon",
			callback:    commandCatchWrapper,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the specified Pokemon's details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
