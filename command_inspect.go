package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify a Pokémon to inspect")
	}
	originalName := args[0]
	pokemonName := strings.ToLower(originalName)

	if pokemon, exists := cfg.pokedex[pokemonName]; exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		// Print types
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.TypeName)
		}
	} else {
		fmt.Println("you have not caught that pokemon")

	}
	return nil
}
