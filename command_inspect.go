package main

import "fmt"

func commandInspect(c *Config) error {
	name := c.inspectpokemon
	pokemon, ok := caughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught the pokemon")
		return nil
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Println("  -", tp.Type.Name)
	}

	return nil
}
