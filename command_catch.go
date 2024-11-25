package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(cfg *config, r *rand.Rand, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify a Pokemon to catch")
	}
	originalName := args[0]
	pokemonName := strings.ToLower(originalName)
	if _, exists := cfg.pokedex[pokemonName]; exists {
		return fmt.Errorf("you've already caught %s", originalName)
	}
	pokemonURL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName + "/"

	responseByte, err := getHTTP(pokemonURL, cfg.cache)
	if err != nil {
		if strings.Contains(err.Error(), "unexpected status code: 404") {
			return fmt.Errorf("could not find the specified Pokémon, please check the name")
		} else {
			return fmt.Errorf("failed to fetch Pokemon data: %w", err)
		}
	}

	var pokemonData PokemonDetail
	err = json.Unmarshal(responseByte, &pokemonData)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	baseExp := pokemonData.Exp
	fmt.Printf("Throwing a Pokeball at %s...\n", originalName)
	catch := CatchPokemon(r, baseExp)
	if catch {
		cfg.pokedex[pokemonName] = pokemonData
		fmt.Printf("%s was caught!\n", originalName)
	} else {
		fmt.Printf("%s has escaped!\n", originalName)
	}
	return nil
}

func CatchPokemon(r *rand.Rand, baseExp int) bool {
	baseChance := 70.0
	adjustedExp := float64(baseExp) / 100.0
	successRate := baseChance / (adjustedExp + 1)

	if successRate < 0 {
		successRate = 10
	}
	fmt.Printf("Success Rate: %f\n", successRate)
	randomChance := float64(r.Intn(100))
	fmt.Printf("Random Chance: %f\n", randomChance)
	return randomChance < successRate
}
