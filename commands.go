package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

// nextLocations needs to take in a object that holds the URL
func nextLocations(config *Config) error {
	URL := baseURL + "/location-area"
	
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

	// all location areas
	var batchList NamedAPIResourceList
	if err = json.Unmarshal(bytes, &batchList); err != nil {
		return err
	}

	
	config.next = batchList.Next
	config.previous = batchList.Previous

	for _, result := range batchList.Results {
		fmt.Println(result.Name)
	    //batchList.Next = &config.next
	}
	
	return nil 
}