package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
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
		"catch": {
			name: "catch",
			description: "Allows an attempt to catch a pokemon",
			callback: catchPokemon,
		},
	}
}

// Exits the pokedex
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

// Help for the user, displays commands
func commandHelp(config *Config, name string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Ϟ(๑⚈ ․̫ ⚈๑)⋆")
	fmt.Println("\t\tUsage!")
	fmt.Println()
	for _, cmd := range commands {
		
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		//fmt.Println()
	}
	fmt.Println()
	return nil
}

// Next 20 location areas with a cache being passed in
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

// Previous 20 locations
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

// Lists all pokemon in current location. First use of name, parameter two
func exploreLocation(config *Config, name string) error {
	fmt.Printf("Exploring %s area...", name)
	// 1: Check cache if name exists in cache, if not attempt to fetch from the URL. I have decided that only this function edits "name". First name will never exist
	startURL := "https://pokeapi.co/api/v2/location-area/" + name
	data, exist := config.cache.Get(startURL)

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


// Attempt to catch a pokemon, second use of name
func catchPokemon(config *Config, name string) error {
	// empty name request
	if name == "" {
		fmt.Println("The name of the pokemon cannot be blank!")
		return nil
	}

	pokemonName := name
	_, ok := config.pokedex[pokemonName]
	if ok {
		fmt.Println("You already caught this pokemon!")
		return nil
	}
	startURL := baseURL + "/pokemon/" + name + "/"
	res, err := http.Get(startURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
						// Pokemon does not exist in the api
	if res.StatusCode >= 400 {
		return fmt.Errorf("Woops! Something went wrong! Check the pokemon name! StatusCode: %d", res.StatusCode)
	}
	
	// JSON processing
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// Pokemon Object + Usage
	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}
	chance := rand.Intn(3)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	// user has a 1 in 3 chance of catching it (for now)
	if chance == 1 {
							// full pokemon object for now. Slim it down later. We need the ID and the Name
		config.pokedex[name] = pokemon
		fmt.Printf("Congratulations! %s has been caught and added to the pokedex!\n", name)
		return nil
	}

	fmt.Printf("Oh No! %s ran away!\n", name)
	return nil

}


