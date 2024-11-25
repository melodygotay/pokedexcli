package main

import (
	"encoding/json"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("please specify a location area name to explore")
	}
	locationName := args[0]

	// Construct the API URL using the location name
	apiURL := baseURL + locationName

	// Fetch the response from the API, utilizing caching
	responseBytes, err := getHTTP(apiURL, cfg.cache)
	if err != nil {
		return fmt.Errorf("failed to fetch location data: %w", err)
	}

	var detailedLocationResp DetailedLocationAreaResp

	// Parse the JSON response into the detailed location structure
	err = json.Unmarshal(responseBytes, &detailedLocationResp)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// Extract Pokémon names from the API response
	pokemonNames := make([]string, len(detailedLocationResp.PokemonEncounters))
	for i, encounter := range detailedLocationResp.PokemonEncounters {
		pokemonNames[i] = encounter.Pokemon.Name
	}

	// Display the found Pokémon names
	fmt.Printf("Exploring %s...", locationName)
	fmt.Println("\nFound Pokemon:")
	for _, name := range pokemonNames {
		fmt.Println("-", name)
	}

	return nil
}
