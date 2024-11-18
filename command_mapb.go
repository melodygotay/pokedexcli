package main

import (
	"fmt"
)

func commandMapB(cfg *config) error {
	var current string

	if cfg.Previous == "" {
		current = baseURL
	} else {
		current = cfg.Previous
	}

	resp, err := getHTTP(current)
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
