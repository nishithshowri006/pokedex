package main

import (
	"fmt"
	"math/rand"
)

type Pokemon struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	BaseExperience int `json:"base_experience"`
}

func commandCatch(c *Config) error {
	pokemonName := c.pokemonName
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	baseUrl := BASEURL + "pokemon/" + pokemonName + "/"
	pokemon, err := pokedexCatchRequest(baseUrl)
	if err != nil {
		return err
	}
	if rand.Intn(pokemon.BaseExperience+1) > int(float32(pokemon.BaseExperience)/1.2) {
		fmt.Printf("%s has escaped!\n", pokemon.Name)
		return nil
	}
	caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}
