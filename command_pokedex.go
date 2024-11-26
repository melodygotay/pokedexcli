package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokemon := cfg.pokedex

	if len(args) > 0 {
		return fmt.Errorf("'pokedex' command does not accept any arguments")
	}

	if len(pokemon) > 0 {
		fmt.Println("Your Pokedex:")
		for _, mons := range pokemon {
			fmt.Printf(" - %s\n", mons.Name)
		}
	} else {
		fmt.Println("you have not caught any pokemon")
	}
	return nil
}
