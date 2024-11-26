package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/location-area/"

func getHTTP(url string, cache *pokecache.Cache) ([]byte, error) {
	// Check cache first
	if cachedData, ok := cache.Get(url); ok {
		//fmt.Println("Cache hit!")
		// Return the cached data directly
		return cachedData, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Store in cache
	cache.Add(url, body)

	// Return the raw JSON bytes
	return body, nil
}

func getLocationAreaResp(url string, cache *pokecache.Cache) (*LocationAreaResp, error) {
	body, err := getHTTP(url, cache)
	if err != nil {
		return nil, err
	}

	var resp LocationAreaResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON from URL %s: %w", url, err)
	}

	return &resp, nil
}

type LocationAreaResp struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type DetailedLocationAreaResp struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type EncounterDetail struct {
	MinLevel int `json:"min_level"`
	MaxLevel int `json:"max_level"`
	Chance   int `json:"chance"`
}

type VersionDetail struct {
	Version struct {
		Name string `json:"name"`
	} `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon         `json:"pokemon"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type PokemonDetail struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Exp    int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []Stat `json:"stats"`
	Types  []Type `json:"types"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Type struct {
		TypeName string `json:"name"`
	} `json:"type"`
}

type config struct {
	Previous string
	Next     string
	cache    *pokecache.Cache
	pokedex  map[string]PokemonDetail
}

func NewConfig() *config {
	return &config{
		cache:    pokecache.NewCache(5 * time.Minute),
		Next:     "",
		Previous: "",
		pokedex:  make(map[string]PokemonDetail),
	}
}
