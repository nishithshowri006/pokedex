package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nishithshowri006/pokedex/internal/pokecache"
)

var cache *pokecache.Cache
var BASEURL = "https://pokeapi.co/api/v2/"
var caughtPokemon = make(map[string]Pokemon)

type Config struct {
	Next           string `json:"next"`
	Previous       string `json:"previous"`
	location       string
	pokemonName    string
	inspectpokemon string
}
type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func startRepl(c *pokecache.Cache) {
	var conf = Config{}
	cache = c
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput(scanner.Text())
		if cmd, ok := getCommands()[inputs[0]]; ok {
			switch cmd.name {
			case "explore":
				if len(inputs) < 2 {
					fmt.Println("explore command needs additional input")
					continue
				}
				conf.location = inputs[1]
			case "catch":

				if len(inputs) < 2 {
					fmt.Println("catch command needs additional input")
					continue
				}
				conf.pokemonName = inputs[1]

			case "inspect":
				if len(inputs) < 2 {
					fmt.Println("inspect command need additional input")
				}
				conf.inspectpokemon = inputs[1]
			}
			err := cmd.callback(&conf)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows 20 location areas in the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous 20 location areas in the pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the Pokemon available in the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throwing ball to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Get information about the caught pokemon",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemons you caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	vals := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return vals
}
