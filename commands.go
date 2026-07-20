package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	//"errors"
)

type cliCommand struct {
	name string
	description string
	callback func(config *Config) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name: "help",
		description: "Print user commands",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Get next 20 locations from the API",
		callback: nextLocations,
	},
	"mapb": {
		name: "mapb",
		description: "Get previous 20 locations from the API",
		callback: previousLocations,
	},
}


func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	func() {
		char := `
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
⠀⠀⠀⠀⠀⠀⠀⠿⠿⠿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠛⠀⠀⠀⠀⠀⠀⠀
				`
		fmt.Println(char)
	}()
	os.Exit(0)

	// never reaches, but its fine for structure
	return nil

}


func commandHelp(config *Config) error {
	func() {
		pikachu := ` 
		Ϟ(๑⚈ ․̫ ⚈๑)⋆
					`
		fmt.Println(pikachu)
	}()
	
	fmt.Println("Welcome to the Pokedex!\nUsage: \n\nhelp: Displays a help message\nexit: Exit the Pokedex")

	return nil

}


func nextLocations(config *Config) error {
	// can add -v for the user of nextLocations later with a query (limit=x)
	URL := baseURL + "/location-area?limit=300"

	if config.next != nil {
		URL = *config.next
	}

	// pull all of the locations
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// raw -> json
	var batchList NamedAPIResourceList
	if err = json.Unmarshal(bytes, &batchList); err != nil {
		return err
	}
	// last page
	if batchList.Next == nil{
		config.previous = batchList.Previous
		for _, result := range batchList.Results {
			fmt.Println(result.Name)
		}

		fmt.Println("\nNOTE: you're on the last page")
		return nil
	}

	// state
	config.next = batchList.Next
	config.previous = batchList.Previous

	for _, result := range batchList.Results {
		fmt.Println(result.Name)
	}

	return nil 

}

func previousLocations(config *Config) error {
	// can add -v for the user of previousLocations later with a query (limit=x)
	URL := baseURL + "/location-area&limit=300"

	if config.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	URL = *config.previous

	// pull 20 of the previous locations 
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// raw bytes -> json
	var batchList NamedAPIResourceList
	if err = json.Unmarshal(bytes, &batchList); err != nil {
		return err
	}

	// state
	config.next = batchList.Next
	config.previous = batchList.Previous // next mapb after this call

	for _, result := range batchList.Results {
		fmt.Println(result.Name)
	}

	
	

	return nil

}
