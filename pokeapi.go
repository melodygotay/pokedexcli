package main

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2/location-area"

func getHTTP(url string) (*LocationAreaResp, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp LocationAreaResp
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resp)
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
}
