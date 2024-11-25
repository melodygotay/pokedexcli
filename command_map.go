package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	var current string

	if cfg.Next == "" {
		current = baseURL
	} else {
		current = cfg.Next
	}

	resp, err := getLocationAreaResp(current, cfg.cache)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	fmt.Println("Locations:")
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	var current string

	if cfg.Previous == "" {
		fmt.Println("You're already on the first page!")
		return nil
	} else {
		current = cfg.Previous
	}

	resp, err := getLocationAreaResp(current, cfg.cache)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	fmt.Println("Locations:")
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
