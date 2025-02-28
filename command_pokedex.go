package main

import "fmt"

func commandPokedex(c *Config) error {
	if len(caughtPokemon) < 1 {
		fmt.Println("Empty pokedex")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range caughtPokemon {
		fmt.Println(" -", caughtPokemon[key].Name)
	}
	return nil
}
