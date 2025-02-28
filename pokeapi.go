package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func pokedexRequest(baseURL string) (Config, error) {
	var body []byte
	var config Config
	var unwrap = struct {
		Count int
		Config
		Results []LocationUrl
	}{}

	if data, ok := cache.Get(baseURL); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			return config, nil
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return config, nil
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return config, nil
		}
		cache.Add(baseURL, body)
	}
	if err := json.Unmarshal(body, &unwrap); err != nil {
		return config, nil
	}
	for _, res := range unwrap.Results {
		fmt.Println(res.Name)
	}
	config = unwrap.Config
	return config, nil
}

func pokedexLocationRequest(baseURL string) error {
	var body []byte
	var pokemons struct {
		PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
	}
	if data, ok := cache.Get(baseURL); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(baseURL, body)
	}
	if err := json.Unmarshal(body, &pokemons); err != nil {
		return err
	}
	for _, res := range pokemons.PokemonEncounters {
		fmt.Println(res.Pokemon.Name)
	}
	return nil
}

func pokedexCatchRequest(baseURL string) (pokemon Pokemon, err error) {
	var body []byte
	if data, ok := cache.Get(baseURL); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", baseURL, nil)
		if err != nil {
			return pokemon, err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("here1")
			return pokemon, err
		}
		if res.StatusCode != 200 {
			return pokemon, fmt.Errorf("Pokemon doesn't exist")
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return pokemon, err
		}
		cache.Add(baseURL, body)
	}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return pokemon, err
	}
	return pokemon, err
}
