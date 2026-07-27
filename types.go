package main

const baseURL = "https://pokeapi.co/api/v2"

// Pagination
type NamedAPIResourceList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// data for pokemon in a location area
type LocationAreaPokemon struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		// this will be used later perhaps
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// The higher the base experience the higher it is to catch 

// Pokedex displays name, etc. Can be very detailed if you like 
type Pokemon struct {
    Name           string `json:"name"`
    BaseExperience int    `json:"base_experience"`
    Height         int    `json:"height"`
    Weight         int    `json:"weight"`
    Stats          []struct {
        BaseStat int `json:"base_stat"`
        Stat     struct{ Name string `json:"name"` } `json:"stat"`
    } `json:"stats"`
    Types          []struct {
        Type struct{ Name string `json:"name"` } `json:"type"`
    } `json:"types"`
}