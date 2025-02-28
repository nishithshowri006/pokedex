package main

import "fmt"

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}

func commandExplore(c *Config) error {
	baseURL := BASEURL + "location-area/" + c.location
	if err := pokedexLocationRequest(baseURL); err != nil {
		fmt.Println("Encountered error: %w", err)
		return err
	}
	return nil
}
