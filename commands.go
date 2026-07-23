package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, name string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print user commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next 20 locations from the API",
			callback:    nextLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 locations from the API",
			callback:    previousLocations,
		},
		"explore": {
			name: "explore",
			description: "Shows a list of all the pokemon located in the specified location",
			callback: exploreLocation,
		},
	}
}

func commandExit(config *Config, name string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println(`
⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⠰⣾⣿⣶⣶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣠⣴⣶⣶⣿⠁⠀⠙⠿⣿⣿⣿⡟⠶⠀⠀⠀⠀⠀⢲⣄⠀⠀⠀⠀⠀
⠀⢀⣶⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀⢻⣿⣧⠀⠀⠀⠀⠀⢀⣾⣿⣿⣦⡀⠀⠀
⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣦⣄⡀⠀⢸⣿⣿⡄⠀⠀⢀⣠⣿⣿⣿⣿⣿⣿⣦⠀
⣿⡟⠉⠀⠀⠈⢻⣿⣿⠿⢿⣿⣿⣷⣾⣿⣿⣷⣶⣾⣿⣿⣿⣿⣿⡿⢿⣿⣿⣇
⢻⠁⠀⠀⠀⠀⠀⠉⣠⣶⡟⠛⣿⣿⣿⣿⣿⣿⣿⠟⠛⣿⣿⣿⠋⠀⣾⣟⢻⣿
⠀⠀⠀⠀⠀⠀⠀⠈⠻⠙⢇⣼⣿⣿⣿⣿⣿⣿⣿⡆⠀⠟⠻⠟⠀⠀⣾⣿⠈⡟
⠀⠀⠀⠀⠀⣀⣤⣴⣤⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⢸⣿⡿⠇⠀
⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⢀⣀⡤⠊⠀⠀⠀
⠀⠀⠀⠘⠿⣿⣿⣿⣿⣿⣿⡿⠿⠿⠛⠛⠻⠿⠿⢿⣿⣿⣿⠛⠋⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠿⠿⠿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠛⠀⠀⠀⠀⠀⠀⠀`)
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, name string) error {
	fmt.Println("\n\t\tϞ(๑⚈ ․̫ ⚈๑)⋆")
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}


func nextLocations(config *Config, name string) error {
	if config.next == nil {
		fmt.Println("you're on the last page")
		return nil
	}

	URL := *config.next

	// check the cache first
	data, exist := config.cache.Get(URL)
	if !exist {
		res, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		config.cache.Add(URL, data)
	}

	var batchList NamedAPIResourceList
	if err := json.Unmarshal(data, &batchList); err != nil {
		return err
	}

	config.next = batchList.Next
	config.previous = batchList.Previous

	for _, result := range batchList.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func previousLocations(config *Config, name string) error {
	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	URL := *config.previous

	// check the cache first
	data, exist := config.cache.Get(URL)
	if !exist {
		res, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		config.cache.Add(URL, data)
	}

	var batchList NamedAPIResourceList
	if err := json.Unmarshal(data, &batchList); err != nil {
		return err
	}

	config.next = batchList.Next
	config.previous = batchList.Previous

	for _, result := range batchList.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Lists all pokemon in current location. First use of name
func exploreLocation(config *Config, name string) error {
	fmt.Println("Exploring pastoria-city-area...")

	// 1: Check cache if name exists in cache, if not attempt to fetch from the URL. I have decided that only this function edits "name". First name will never exist
	startURL := "https://pokeapi.co/api/v2/location-area/" + name
	data, exist := config.cache.Get(startURL)

	// 2: Download new data (non-exist)
	if !exist {
		res, err :=  http.Get(startURL) // we are now fetching a specific area, not 20 areas
		if err != nil {
			return err
		}
		defer res.Body.Close()
		// possibly an invalid name parameter was passed 
		if res.StatusCode >= 400 {
			return fmt.Errorf("fetching %s: status %d", startURL, res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		config.cache.Add(startURL, data) // raw bytes still like always
	}

	var locationAreaPokemon LocationAreaPokemon
	if err := json.Unmarshal(data, &locationAreaPokemon); err != nil {
		return err
	}

	for _, encounter := range locationAreaPokemon.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
	
	}

	


