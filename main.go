package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/fanatcx/pokedexcli/internal/pokecache"
)

// Config stores next address, previous address and cache data for location. Next = next 20 location areas. Previous = previous 20 areas. Previous begins at nil.
// Dynamic cache for faster loading.

type Config struct {
	name     string
	next     *string
	previous *string
	pokedex  map[string]Pokemon
	cache    *pokecache.Cache
	
}

func main() {
	startURL := baseURL + "/location-area/"
	pokedex := make(map[string]Pokemon)

	cfg := &Config{
		name: "",
		next:  &startURL,
		pokedex: pokedex,
		cache: pokecache.NewCache(5 * time.Second),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		cleaned := CleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}
		// User submits one command, which is cleaned of whitespace
		comm, exists := commands[cleaned[0]] 
		if !exists {
			fmt.Println("Invalid command")
			continue
		}
		
		// User submits exactly two words. We dont have a way to check if its proper yet
		if len(cleaned) == 2 {
			cfg.name = cleaned[1]
		}

		// pass the Config object as a pointer, name is the second argument for specific functions
		if err := comm.callback(cfg, cfg.name); err != nil {
			fmt.Println(err.Error())
			continue
		}
		cfg.name = "" // reset to nil for next loop 
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}