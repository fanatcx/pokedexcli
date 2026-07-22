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
	callback    func(config *Config) error
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
	}
}

func commandExit(config *Config) error {
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

func commandHelp(config *Config) error {
	fmt.Println("\n\t\tϞ(๑⚈ ․̫ ⚈๑)⋆")
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}


func nextLocations(config *Config) error {
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

func previousLocations(config *Config) error {
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