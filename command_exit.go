package main

import (
	"fmt"
	"os"
)

func commandExit(_ *Config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
