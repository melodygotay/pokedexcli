package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/location-area"

func getHTTP(url string, cache *pokecache.Cache) (*LocationAreaResp, error) {
	// Check cache first
	if cachedData, ok := cache.Get(url); ok {
		var resp LocationAreaResp
		err := json.Unmarshal(cachedData, &resp)
		if err != nil {
			return nil, err
		}
		fmt.Println("Cache hit!")
		return &resp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Store in cache
	cache.Add(url, body)

	var resp LocationAreaResp
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
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

type config struct {
	Previous string
	Next     string
	cache    *pokecache.Cache
}

func NewConfig() *config {
	return &config{
		cache:    pokecache.NewCache(5 * time.Minute),
		Next:     "",
		Previous: "",
	}
}
