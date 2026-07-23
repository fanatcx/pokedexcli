package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/fanatcx/pokedexcli/internal/pokecache"
	
)

type Config struct {
	name     string
	next     *string
	previous *string
	cache    *pokecache.Cache
}

func main() {
	startURL := baseURL + "/location-area/"
	

	cfg := &Config{
		name: "",
		next:  &startURL,
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

		comm, exists := commands[cleaned[0]]
		if !exists {
			fmt.Println("Invalid command")
			continue
		}
		
		// limit all commands to two params as the max
		if len(cleaned) > 1 {
			cfg.name = cleaned[1]
		}

		// either pass the original address or pass a dereferenced name
		if err := comm.callback(cfg, cfg.name); err != nil {
			fmt.Println(err.Error())
			continue
		}
		cfg.name = "" // only one function takes name and does something useful, and im ok with it
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}