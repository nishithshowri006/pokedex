package main

var counter int

type LocationUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func commandMap(conf *Config) error {
	baseURL := BASEURL + "location-area/"
	if conf.Next != "" {
		baseURL = conf.Next
	}
	ex, err := pokedexRequest(baseURL)
	if err != nil {
		return err
	}
	conf.Previous = ex.Previous
	conf.Next = ex.Next
	return nil
}
