package main

import "fmt"

func commandMapB(conf *Config) error {
	if conf == nil || conf.Previous == "" {
		fmt.Println("you’re on the first page")
		return nil
	}
	baseURL := conf.Previous
	ex, err := pokedexRequest(baseURL)
	if err != nil {
		return err
	}
	conf.Previous = ex.Previous
	conf.Next = ex.Next
	return nil
}
